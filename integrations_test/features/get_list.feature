# file: features/get_list.feature

Feature: Server grpc get list object

  Scenario: Get event list week
    When I GetEventList with type week gprc-request to "server:8000"
    Then The response must contain event list on week

  Scenario: Get event list month
    When I GetEventList with type month gprc-request to "server:8000"
    Then The response must contain event list on month

  Scenario: Get event list year
    When I GetEventList with type year gprc-request to "server:8000"
    Then The response must contain event list on year