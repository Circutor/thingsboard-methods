Feature: Telemetry Controller
  Background:
    Given I open a ThingsBoard session
    And I init controller telemetry

  Scenario: Successful outcome. Method Get last TimeSeries.
    When I wish the entity "DEVICE" and entity ID "2d0f8500-bf9b-11eb-8d1e-ab995e7e469b"