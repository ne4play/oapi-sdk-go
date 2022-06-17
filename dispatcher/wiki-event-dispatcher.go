// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/feishu/oapi-sdk-go/service/wiki/v2"
)

func (dispatcher *EventReqDispatcher) SpaceNodeMoveDocsToWikiApplyHandledV2(handler func(ctx context.Context, event *wiki.SpaceNodeMoveDocsToWikiApplyHandledEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["wiki.space.node.move_docs_to_wiki_apply_handled_v2"] = wiki.NewSpaceNodeMoveDocsToWikiApplyHandledEventHandler(handler)
	return dispatcher
}