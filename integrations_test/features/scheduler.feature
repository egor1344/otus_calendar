# file: features/scheduler.feature

Feature: Scheduler

  Scenario: Send mailing and Clear old events
    When Run scheduler
    Then Exists events update send True and Clear old events

