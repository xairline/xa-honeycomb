use serde::{Deserialize, Serialize};
use serde_yaml;
use std::collections::HashMap;
use std::fs;
use std::path::Path;

// Command struct
#[derive(Debug, Serialize, Deserialize)]
pub struct Command {
    #[serde(rename = "command_str", skip_serializing_if = "Option::is_none")]
    pub command_str: Option<String>,
    pub command: Option<serde_json::Value>,
}

// Dataref struct
#[derive(Debug, Serialize, Deserialize)]
pub struct Dataref {
    #[serde(rename = "dataref_str", skip_serializing_if = "Option::is_none")]
    pub dataref_str: Option<String>,
    pub dataref: Option<serde_json::Value>,
    pub index: Option<i32>,
}

// DatarefCondition struct
#[derive(Debug, Serialize, Deserialize)]
pub struct DatarefCondition {
    #[serde(rename = "dataref_str", skip_serializing_if = "Option::is_none")]
    pub dataref_str: Option<String>,
    pub dataref: Option<serde_json::Value>,
    pub index: Option<i32>,
    pub operator: Option<String>,
    pub threshold: Option<f32>,
    pub expr: Option<String>, // Simplified vm.Program to String for now
    pub env: Option<HashMap<String, serde_json::Value>>,
}

// Metadata struct
#[derive(Debug, Serialize, Deserialize)]
pub struct Metadata {
    pub name: Option<String>,
    pub description: Option<String>,
    pub selectors: Option<Vec<String>>,
}

// Profiles
#[derive(Debug, Serialize, Deserialize)]
pub struct ConditionProfile {
    pub datarefs: Option<Vec<DatarefCondition>>,
    pub condition: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DatarefProfile {
    pub datarefs: Option<Vec<Dataref>>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct LEDProfile {
    #[serde(flatten)]
    pub condition_profile: ConditionProfile,
    #[serde(skip_serializing, skip_deserializing)]
    pub on: Option<fn()>,
    #[serde(skip_serializing, skip_deserializing)]
    pub off: Option<fn()>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DataProfile {
    #[serde(flatten)]
    pub dataref_profile: DatarefProfile,
    pub value: Option<f32>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct KnobProfile {
    #[serde(flatten)]
    pub dataref_profile: DatarefProfile,
    pub commands: Option<Vec<Command>>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ButtonProfile {
    #[serde(rename = "single_click")]
    pub single_click: Option<Vec<Command>>,
    #[serde(rename = "double_click")]
    pub double_click: Option<Vec<Command>>,
}

// Composite structs
#[derive(Debug, Serialize, Deserialize)]
pub struct Knobs {
    pub ap_hdg: Option<KnobProfile>,
    pub ap_vs: Option<KnobProfile>,
    pub ap_alt: Option<KnobProfile>,
    pub ap_ias: Option<KnobProfile>,
    pub ap_crs: Option<KnobProfile>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Leds {
    pub hdg: Option<LEDProfile>,
    pub nav: Option<LEDProfile>,
    pub alt: Option<LEDProfile>,
    pub apr: Option<LEDProfile>,
    pub vs: Option<LEDProfile>,
    pub ap: Option<LEDProfile>,
    pub ias: Option<LEDProfile>,
    pub rev: Option<LEDProfile>,
    pub gear: Option<LEDProfile>,
    pub master_warn: Option<LEDProfile>,
    pub master_caution: Option<LEDProfile>,
    pub fire: Option<LEDProfile>,
    pub oil_low_pressure: Option<LEDProfile>,
    pub fuel_low_pressure: Option<LEDProfile>,
    pub anti_ice: Option<LEDProfile>,
    pub eng_starter: Option<LEDProfile>,
    pub apu: Option<LEDProfile>,
    pub vacuum: Option<LEDProfile>,
    pub hydro_low_pressure: Option<LEDProfile>,
    pub aux_fuel_pump: Option<LEDProfile>,
    pub parking_brake: Option<LEDProfile>,
    pub volt_low: Option<LEDProfile>,
    pub doors: Option<LEDProfile>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Data {
    pub ap_state: Option<DataProfile>,
    pub ap_alt_step: Option<DataProfile>,
    pub ap_vs_step: Option<DataProfile>,
    pub ap_ias_step: Option<DataProfile>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Conditions {
    pub bus_voltage: Option<ConditionProfile>,
    pub retractable_gear: Option<ConditionProfile>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Buttons {
    pub hdg: Option<ButtonProfile>,
    pub nav: Option<ButtonProfile>,
    pub alt: Option<ButtonProfile>,
    pub apr: Option<ButtonProfile>,
    pub vs: Option<ButtonProfile>,
    pub ap: Option<ButtonProfile>,
    pub ias: Option<ButtonProfile>,
    pub rev: Option<ButtonProfile>,
}

// Final Profile struct
#[derive(Debug, Serialize, Deserialize)]
pub struct Profile {
    pub metadata: Option<Metadata>,
    pub buttons: Option<Buttons>,
    pub knobs: Option<Knobs>,
    pub leds: Option<Leds>,
    pub data: Option<Data>,
    pub conditions: Option<Conditions>,
}
