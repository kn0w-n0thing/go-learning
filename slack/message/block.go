package message

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
)

func BuildLeaveRequestMessageOptions() slack.MsgOption {
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", "*Leave request*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Sections
	nameField := slack.NewTextBlockObject("mrkdwn", "*Applicant:*\nZHANG San", false, false)
	departmentField := slack.NewTextBlockObject("mrkdwn", "*Department:*\nRD Center", false, false)
	personalInfoSlice := make([]*slack.TextBlockObject, 0)
	personalInfoSlice = append(personalInfoSlice, nameField)
	personalInfoSlice = append(personalInfoSlice, departmentField)
	personalInfoSection := slack.NewSectionBlock(nil, personalInfoSlice, nil)

	leavePeriodField := slack.NewTextBlockObject("mrkdwn", "*Period:*\nAug 10 - Aug 13", false, false)
	leaveDurationField := slack.NewTextBlockObject("mrkdwn", "*Duration:*\n4 days", false, false)
	leaveInfoSlice := make([]*slack.TextBlockObject, 0)
	leaveInfoSlice = append(leaveInfoSlice, leavePeriodField)
	leaveInfoSlice = append(leaveInfoSlice, leaveDurationField)
	leaveInfoSection := slack.NewSectionBlock(nil, leaveInfoSlice, nil)

	reasonFiled := slack.NewTextBlockObject("mrkdwn", "*Reason:*\nGoing to hospital.", false, false)
	reasonSlice := make([]*slack.TextBlockObject, 0)
	reasonSlice = append(reasonSlice, reasonFiled)
	reasonSection := slack.NewSectionBlock(nil, reasonSlice, nil)

	// Approve and Deny Buttons
	approveButtonText := slack.NewTextBlockObject("plain_text", "Approve", false, false)
	approveButton := slack.ButtonBlockElement{
		Type:  slack.METButton,
		Text:  approveButtonText,
		Value: "reject_button",
		Style: slack.StylePrimary,
	}

	rejectButtonText := slack.NewTextBlockObject("plain_text", "Reject", false, false)
	rejectButton := slack.ButtonBlockElement{
		Type:  slack.METButton,
		Text:  rejectButtonText,
		Value: "reject_button",
		Style: slack.StyleDanger,
	}

	actionBlock := slack.NewActionBlock("", approveButton, rejectButton)

	// Build Message with blocks created above

	msg := slack.NewBlockMessage(
		headerSection,
		personalInfoSection,
		leaveInfoSection,
		reasonSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	return slack.MsgOptionBlocks(
		headerSection,
		personalInfoSection,
		leaveInfoSection,
		reasonSection,
		actionBlock,
	)
}
