# How to read/change profiles

## Sample Profile

```yaml
# Metadata about the profile
metadata:
  # The name of the profile
  name: "Sample Profile"
  # A description of the profile
  description: "This is a sample configuration profile."
  # A list of selectors to apply the profile
  selectors:
    - "default"
    - "custom"

# Knob configurations
knobs:
  # Configuration for the Autopilot Heading (AP_HDG) knob
  ap_hdg:
    # List of data references associated with this knob
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/heading"  # Data reference string
        index: 0                                       # Index (if applicable)
    # List of commands associated with this knob
    commands:
      - command_str: "sim/autopilot/heading_up"        # Command to increase heading
      - command_str: "sim/autopilot/heading_down"      # Command to decrease heading

  # Configuration for the Autopilot Vertical Speed (AP_VS) knob
  ap_vs:
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/vertical_speed"
        index: 0
    commands:
      - command_str: "sim/autopilot/vertical_speed_up"
      - command_str: "sim/autopilot/vertical_speed_down"

# LED configurations
leds:
  # Configuration for the Heading (HDG) LED
  hdg:
    # Data references used to determine LED state
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/heading_mode"  # Data reference string
        index: 0
        operator: "=="                                     # Operator for comparison (e.g., ==, !=, >, <)
        threshold: 1.0                                     # Threshold value for the condition
    # Condition to evaluate (can use expressions)
    condition: "datarefs[0] == 1"

  # Configuration for the Navigation (NAV) LED
  nav:
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/nav_mode"
        index: 0
        operator: "=="
        threshold: 1.0
    condition: "datarefs[0] == 1"

# Data configurations
data:
  # Configuration for the Autopilot State
  ap_state:
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/autopilot_state"  # Data reference string
        index: 0
    # Value to set or monitor (optional)
    value: 1.0

  # Configuration for the Autopilot Altitude Step
  ap_alt_step:
    datarefs:
      - dataref_str: "sim/cockpit/autopilot/altitude_step"
        index: 0
    value: 100.0  # Step value for altitude adjustments

# Condition configurations
conditions:
  # Condition for Bus Voltage
  bus_voltage:
    datarefs:
      - dataref_str: "sim/cockpit/electrical/bus_voltage"  # Data reference string
        index: 0
        operator: ">"                                      # Operator for comparison
        threshold: 24.0                                     # Threshold voltage
    condition: "datarefs[0] > 24.0"

  # Condition for Retractable Gear
  retractable_gear:
    datarefs:
      - dataref_str: "sim/aircraft/gear/retractable"  # Data reference string
        index: 0
        operator: "=="                               # Operator for comparison
        threshold: 1.0                               # Threshold value (1.0 for true)
    condition: "datarefs[0] == 1"

# Additional configuration sections can be added below following the same structure.
```
## Metadata

## Knobs

## LEDs

## Data

## Conditions

## Buttons (Future)