package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
)

type IPlennerService interface {
	AddMeeting(meeting entities.Meeting)
}
