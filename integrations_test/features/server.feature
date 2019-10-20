# file: features/server.feature

# http://server:8000/

Feature: Server grpc
	As grpc client of event service
	In order to understand that the user was informed about registration
	I want to receive event from notifications queue

	Scenario: Add event user
    		When I AddEvent gprc-request to "server:8000"
    		Then The response should match my Event