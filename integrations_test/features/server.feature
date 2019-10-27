# file: features/server.feature

Feature: Server grpc
	As grpc client of event service
	In order to understand that the user was informed about registration
	I want to receive event from notifications queue

	Scenario: Add event user
    		When I AddEvent gprc-request to "server:8000"
    		Then The response should match my Event

	Scenario: Get event user
		When I Get event gprc-request to "server:8000"
		Then The response must contain my event

	Scenario: Update event user
		When I UpdateEvent gprc-request to "server:8000"
		Then The response must contain my update event

	Scenario: Get event list week
		When I GetEventList with type week gprc-request to "server:8000"
		Then The response must contain event list on week

	Scenario: Get event list month
		When I GetEventList with type month gprc-request to "server:8000"
		Then The response must contain event list on month

	Scenario: Get event list year
		When I GetEventList with type year gprc-request to "server:8000"
		Then The response must contain event list on year