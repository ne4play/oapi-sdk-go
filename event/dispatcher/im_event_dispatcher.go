// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/service/im/v1"
)

func (dispatcher *EventReqDispatcher) OnChatDisbandedV1(handler func(ctx context.Context, event *im.ChatDisbandedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.disbanded_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.disbanded_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.disbanded_v1"] = im.NewChatDisbandedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatUpdatedV1(handler func(ctx context.Context, event *im.ChatUpdatedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.updated_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.updated_v1"] = im.NewChatUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatMemberBotAddedV1(handler func(ctx context.Context, event *im.ChatMemberBotAddedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.bot.added_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.bot.added_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.bot.added_v1"] = im.NewChatMemberBotAddedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatMemberBotDeletedV1(handler func(ctx context.Context, event *im.ChatMemberBotDeletedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.bot.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.bot.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.bot.deleted_v1"] = im.NewChatMemberBotDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatMemberUserAddedV1(handler func(ctx context.Context, event *im.ChatMemberUserAddedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.added_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.added_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.added_v1"] = im.NewChatMemberUserAddedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatMemberUserDeletedV1(handler func(ctx context.Context, event *im.ChatMemberUserDeletedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.deleted_v1"] = im.NewChatMemberUserDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnChatMemberUserWithdrawnV1(handler func(ctx context.Context, event *im.ChatMemberUserWithdrawnEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.withdrawn_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.withdrawn_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.withdrawn_v1"] = im.NewChatMemberUserWithdrawnEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMessageMessageReadV1(handler func(ctx context.Context, event *im.MessageMessageReadEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.message_read_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.message_read_v1")
	}
	dispatcher.eventType2EventHandler["im.message.message_read_v1"] = im.NewMessageMessageReadEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMessageReceiveV1(handler func(ctx context.Context, event *im.MessageReceiveEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.receive_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.receive_v1")
	}
	dispatcher.eventType2EventHandler["im.message.receive_v1"] = im.NewMessageReceiveEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMessageReactionCreatedV1(handler func(ctx context.Context, event *im.MessageReactionCreatedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.reaction.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.reaction.created_v1")
	}
	dispatcher.eventType2EventHandler["im.message.reaction.created_v1"] = im.NewMessageReactionCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMessageReactionDeletedV1(handler func(ctx context.Context, event *im.MessageReactionDeletedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.reaction.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.reaction.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.message.reaction.deleted_v1"] = im.NewMessageReactionDeletedEventHandler(handler)
	return dispatcher
}