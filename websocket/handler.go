package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

const (
	UUIDHeaderKey = "X-WebSocket-UUID"
)

const (
	pingPeriod = 30 * time.Second
)

type Op string

const (
	OpSubscribeId      Op = "subscriber-id"
	OpSubscribeTopic      = "subscribe-topic"
	OpBroadcast           = "broadcast"
	OpBrodcasted          = "brodcasted"
	OpUnsubscribeTopic    = "unsubscribe-topic"
)

type message struct {
	Op      Op     `json:"op"`
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
}

type webSocketHandler struct {
	subscribers   map[string]net.Conn
	subscriptions map[string][]string
}

func (handler webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("Error upgrading Connection: %v", err)
		return
	}

	subscriberID := uuid.New().String()
	w.Header().Set(UUIDHeaderKey, subscriberID)

	handler.subscribers[subscriberID] = conn

	connectedMessage := message{Op: OpSubscribeId, Payload: subscriberID}
	err = handler.send(subscriberID, connectedMessage)
	if err != nil {
		log.Printf("error sending subscriber id: %v", err)
		return
	}

	done := make(chan struct{})

	go handler.ping(conn, subscriberID, done)
	handler.read(conn, subscriberID, done)
}

func (handler webSocketHandler) read(conn net.Conn, subscriberID string, done <-chan struct{}) {
	for {
		msg, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Printf("Error reading client data: %v", err)
			break
		}

		clientMessage := message{}
		err = json.Unmarshal(msg, &clientMessage)
		if err != nil {
			log.Printf("Error unpacking json: %v", err)
		}

		handler.processMessage(subscriberID, clientMessage)
	}
}

func (handler webSocketHandler) processMessage(subscriberID string, msg message) error {
	switch msg.Op {
	case OpSubscribeTopic:
		handler.subscribe(subscriberID, msg.Topic)
	case OpBroadcast:
		handler.broadcast(subscriberID, msg.Topic, msg.Payload)
	case OpUnsubscribeTopic:
		handler.unsubscribeTopic(subscriberID, msg.Topic)
	}
	return nil
}

func (handler webSocketHandler) ping(conn net.Conn, subscriberID string, done <-chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := wsutil.WriteServerMessage(conn, ws.OpPing, []byte{})
			if err != nil {
				handler.removeSubscriber(subscriberID)
				log.Printf("Disconnecting client %s: %v", subscriberID, err)
				return
			}
		case <-done:
			return
		}
	}
}

func (handler *webSocketHandler) subscribe(subscriberID string, topicName string) {
	if _, ok := handler.subscriptions[topicName]; !ok {
		handler.subscriptions[topicName] = make([]string, 0)
	}
	handler.subscriptions[topicName] = append(handler.subscriptions[topicName], subscriberID)
}

func (handler *webSocketHandler) unsubscribeTopic(subscriberID string, topicName string) {
	log.Printf("Unsuscribe %s from topic '%s'\n", subscriberID, topicName)
	subscription := make([]string, 0)
	for i, subscriber := range handler.subscriptions[topicName] {
		if subscriber == subscriberID {
			subscription = append(handler.subscriptions[topicName][:i], handler.subscriptions[topicName][i+1:]...)
		}
	}
	handler.subscriptions[topicName] = subscription
}

func (handler *webSocketHandler) broadcast(subscriberID string, topicName string, msg string) error {
	if _, ok := handler.subscriptions[topicName]; ok {
		for subscriberIndex, subscriber := range handler.subscriptions[topicName] {
			if _, ok := handler.subscribers[subscriber]; ok {
				err := handler.send(subscriber, message{Op: OpBrodcasted, Topic: topicName, Payload: msg})
				if err != nil {
					return err
				}
			} else {
				log.Printf("Suscriber %s doesn't exist, deleting subscription", subscriber)
				handler.subscriptions[topicName] = append(
					handler.subscriptions[topicName][:subscriberIndex],
					handler.subscriptions[topicName][subscriberIndex+1:]...)
			}
		}
	}
	return nil
}

func (handler webSocketHandler) send(subscriberID string, msg message) error {
	messageJSON, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Failed to marshal message: %w", err)
	}
	err = wsutil.WriteServerText(handler.subscribers[subscriberID], messageJSON)
	if err != nil {
		return fmt.Errorf("Failed to write message to suscriber %s: %w", subscriberID, err)
	}
	return nil
}

func (handler *webSocketHandler) removeSubscriber(subscriberID string) {
	delete(handler.subscribers, subscriberID)
}

func NewWebSocketHandler() *webSocketHandler {
	handler := &webSocketHandler{
		subscribers:   make(map[string]net.Conn),
		subscriptions: make(map[string][]string),
	}

	return handler
}
