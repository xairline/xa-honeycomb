use crate::plugin_debugln;
use std::fmt::Write;
use std::sync::{Arc, Mutex};

// LED Constants
pub const LED_FUEL_PUMP: u8 = 1;
pub const LED_PARKING_BRAKE: u8 = 2;
pub const LED_LOW_VOLTS: u8 = 4;
pub const LED_DOOR: u8 = 8;

pub const LED_LOW_OIL_PRESS: u8 = 1;
pub const LED_LOW_FUEL_PRESS: u8 = 2;
pub const LED_ANTI_ICE: u8 = 4;
pub const LED_STARTER: u8 = 8;
pub const LED_APU: u8 = 16;
pub const LED_MASTER_CAUTION: u8 = 32;
pub const LED_VACUUM: u8 = 64;
pub const LED_LOW_HYD_PRESS: u8 = 128;

pub const LED_LEFT_GEAR_GREEN: u8 = 1;
pub const LED_LEFT_GEAR_RED: u8 = 2;
pub const LED_NOSE_GEAR_GREEN: u8 = 4;
pub const LED_NOSE_GEAR_RED: u8 = 8;
pub const LED_RIGHT_GEAR_GREEN: u8 = 16;
pub const LED_RIGHT_GEAR_RED: u8 = 32;
pub const LED_MASTER_WARNING: u8 = 64;
pub const LED_ENGINE_FIRE: u8 = 128;

pub const LED_HEADING: u8 = 1;
pub const LED_NAV: u8 = 2;
pub const LED_APR: u8 = 4;
pub const LED_REV: u8 = 8;
pub const LED_ALT: u8 = 16;
pub const LED_VS: u8 = 32;
pub const LED_IAS: u8 = 64;
pub const LED_AP: u8 = 128;

// Shared State
lazy_static::lazy_static! {
    static ref ANUNCIATOR_W2: Mutex<u8> = Mutex::new(0);
    static ref ANUNCIATOR_W1: Mutex<u8> = Mutex::new(0);
    static ref LANDING_GEAR_W: Mutex<u8> = Mutex::new(0);
    static ref AUTO_PILOT_W: Mutex<u8> = Mutex::new(0);
    static ref LED_STATE_CHANGED: Mutex<bool> = Mutex::new(false);
}

// Utility Functions
fn set_bit(val: &mut u8, bit: u8) {
    *val |= bit;
}

fn clear_bit(val: &mut u8, bit: u8) {
    *val &= !bit;
}

fn update_led_state_changed(condition: bool) {
    let mut state_changed = LED_STATE_CHANGED.lock().unwrap();
    *state_changed = *state_changed || condition;
}

// Macro to simplify On/Off function generation
macro_rules! define_led_functions {
    ($name:ident, $bit:expr, $state:ident) => {
        paste::paste! {
            pub fn $name() {
                let mut state = $state.lock().unwrap();
                let before = *state;
                set_bit(&mut *state, $bit);
                update_led_state_changed(before != *state);
            }

            pub fn [< $name _off >]() {
                let mut state = $state.lock().unwrap();
                let before = *state;
                clear_bit(&mut *state, $bit);
                update_led_state_changed(before != *state);
            }
        }
    };
}

// Define functions for ANUNCIATOR_W2 LEDs
define_led_functions!(on_led_fuel_pump, LED_FUEL_PUMP, ANUNCIATOR_W2);
define_led_functions!(on_led_parking_brake, LED_PARKING_BRAKE, ANUNCIATOR_W2);
define_led_functions!(on_led_low_volts, LED_LOW_VOLTS, ANUNCIATOR_W2);
define_led_functions!(on_led_door, LED_DOOR, ANUNCIATOR_W2);

// Define functions for ANUNCIATOR_W1 LEDs
define_led_functions!(on_led_low_oil_press, LED_LOW_OIL_PRESS, ANUNCIATOR_W1);
define_led_functions!(on_led_low_fuel_press, LED_LOW_FUEL_PRESS, ANUNCIATOR_W1);
define_led_functions!(on_led_anti_ice, LED_ANTI_ICE, ANUNCIATOR_W1);
define_led_functions!(on_led_starter, LED_STARTER, ANUNCIATOR_W1);
define_led_functions!(on_led_apu, LED_APU, ANUNCIATOR_W1);
define_led_functions!(on_led_master_caution, LED_MASTER_CAUTION, ANUNCIATOR_W1);
define_led_functions!(on_led_vacuum, LED_VACUUM, ANUNCIATOR_W1);
define_led_functions!(on_led_low_hyd_press, LED_LOW_HYD_PRESS, ANUNCIATOR_W1);

// Define functions for LANDING_GEAR_W LEDs
define_led_functions!(on_led_left_gear_green, LED_LEFT_GEAR_GREEN, LANDING_GEAR_W);
define_led_functions!(on_led_left_gear_red, LED_LEFT_GEAR_RED, LANDING_GEAR_W);
define_led_functions!(on_led_nose_gear_green, LED_NOSE_GEAR_GREEN, LANDING_GEAR_W);
define_led_functions!(on_led_nose_gear_red, LED_NOSE_GEAR_RED, LANDING_GEAR_W);
define_led_functions!(
    on_led_right_gear_green,
    LED_RIGHT_GEAR_GREEN,
    LANDING_GEAR_W
);
define_led_functions!(on_led_right_gear_red, LED_RIGHT_GEAR_RED, LANDING_GEAR_W);

// Define functions for AUTO_PILOT_W LEDs
define_led_functions!(on_led_heading, LED_HEADING, AUTO_PILOT_W);
define_led_functions!(on_led_nav, LED_NAV, AUTO_PILOT_W);
define_led_functions!(on_led_apr, LED_APR, AUTO_PILOT_W);
define_led_functions!(on_led_rev, LED_REV, AUTO_PILOT_W);
define_led_functions!(on_led_alt, LED_ALT, AUTO_PILOT_W);
define_led_functions!(on_led_vs, LED_VS, AUTO_PILOT_W);
define_led_functions!(on_led_ias, LED_IAS, AUTO_PILOT_W);
define_led_functions!(on_led_ap, LED_AP, AUTO_PILOT_W);

// Debugging Function
pub fn debug_print_led_states() {
    let w2 = ANUNCIATOR_W2.lock().unwrap();
    let w1 = ANUNCIATOR_W1.lock().unwrap();
    let lg = LANDING_GEAR_W.lock().unwrap();
    let ap = AUTO_PILOT_W.lock().unwrap();

    let mut output = String::new();

    writeln!(&mut output, "ANUNCIATOR_W2: {:08b}", *w2).unwrap();
    writeln!(&mut output, "ANUNCIATOR_W1: {:08b}", *w1).unwrap();
    writeln!(&mut output, "LANDING_GEAR_W: {:08b}", *lg).unwrap();
    writeln!(&mut output, "AUTO_PILOT_W: {:08b}", *ap).unwrap();

    println!("{}", output);
}

pub fn all_off() {
    let mut w2 = ANUNCIATOR_W2.lock().unwrap();
    let mut w1 = ANUNCIATOR_W1.lock().unwrap();
    let mut lg = LANDING_GEAR_W.lock().unwrap();
    let mut ap = AUTO_PILOT_W.lock().unwrap();
    let mut state_changed = LED_STATE_CHANGED.lock().unwrap();

    *w2 = 0;
    *w1 = 0;
    *lg = 0;
    *ap = 0;
    *state_changed = true;
}

// Simulate sending HID report
fn send_hid_report(hid_report_buffer: &Arc<Mutex<[u8; 5]>>) -> Result<usize, String> {
    let buffer = hid_report_buffer.lock().unwrap();
    // Simulated HID write logic
    plugin_debugln!("Sending HID report: {:?}", *buffer);
    Ok(buffer.len()) // Simulated bytes written
}
