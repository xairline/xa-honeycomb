export namespace pkg {
	
	export class Command {
	    command_str?: string;
	
	    static createFrom(source: any = {}) {
	        return new Command(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.command_str = source["command_str"];
	    }
	}
	export class DatarefCondition {
	    dataref_str?: string;
	    index?: number;
	    operator?: string;
	    threshold?: number;
	
	    static createFrom(source: any = {}) {
	        return new DatarefCondition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dataref_str = source["dataref_str"];
	        this.index = source["index"];
	        this.operator = source["operator"];
	        this.threshold = source["threshold"];
	    }
	}
	export class ConditionProfile {
	    datarefs?: DatarefCondition[];
	    condition?: string;
	
	    static createFrom(source: any = {}) {
	        return new ConditionProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.datarefs = this.convertValues(source["datarefs"], DatarefCondition);
	        this.condition = source["condition"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Conditions {
	    bus_voltage?: ConditionProfile;
	    retractable_gear?: ConditionProfile;
	
	    static createFrom(source: any = {}) {
	        return new Conditions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bus_voltage = this.convertValues(source["bus_voltage"], ConditionProfile);
	        this.retractable_gear = this.convertValues(source["retractable_gear"], ConditionProfile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Dataref {
	    dataref_str?: string;
	    index?: number;
	
	    static createFrom(source: any = {}) {
	        return new Dataref(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dataref_str = source["dataref_str"];
	        this.index = source["index"];
	    }
	}
	export class DataProfile {
	    datarefs?: Dataref[];
	    value?: number;
	
	    static createFrom(source: any = {}) {
	        return new DataProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.datarefs = this.convertValues(source["datarefs"], Dataref);
	        this.value = source["value"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    ap_state?: DataProfile;
	    ap_alt_step?: DataProfile;
	    ap_vs_step?: DataProfile;
	    ap_ias_step?: DataProfile;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ap_state = this.convertValues(source["ap_state"], DataProfile);
	        this.ap_alt_step = this.convertValues(source["ap_alt_step"], DataProfile);
	        this.ap_vs_step = this.convertValues(source["ap_vs_step"], DataProfile);
	        this.ap_ias_step = this.convertValues(source["ap_ias_step"], DataProfile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class KnobProfile {
	    datarefs?: Dataref[];
	    commands?: Command[];
	
	    static createFrom(source: any = {}) {
	        return new KnobProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.datarefs = this.convertValues(source["datarefs"], Dataref);
	        this.commands = this.convertValues(source["commands"], Command);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Knobs {
	    ap_hdg?: KnobProfile;
	    ap_vs?: KnobProfile;
	    ap_alt?: KnobProfile;
	    ap_ias?: KnobProfile;
	    ap_crs?: KnobProfile;
	
	    static createFrom(source: any = {}) {
	        return new Knobs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ap_hdg = this.convertValues(source["ap_hdg"], KnobProfile);
	        this.ap_vs = this.convertValues(source["ap_vs"], KnobProfile);
	        this.ap_alt = this.convertValues(source["ap_alt"], KnobProfile);
	        this.ap_ias = this.convertValues(source["ap_ias"], KnobProfile);
	        this.ap_crs = this.convertValues(source["ap_crs"], KnobProfile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LEDProfile {
	    datarefs?: DatarefCondition[];
	    condition?: string;
	
	    static createFrom(source: any = {}) {
	        return new LEDProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.datarefs = this.convertValues(source["datarefs"], DatarefCondition);
	        this.condition = source["condition"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Leds {
	    hdg?: LEDProfile;
	    nav?: LEDProfile;
	    alt?: LEDProfile;
	    apr?: LEDProfile;
	    vs?: LEDProfile;
	    ap?: LEDProfile;
	    ias?: LEDProfile;
	    rev?: LEDProfile;
	    gear?: LEDProfile;
	    master_warn?: LEDProfile;
	    master_caution?: LEDProfile;
	    fire?: LEDProfile;
	    oil_low_pressure?: LEDProfile;
	    fuel_low_pressure?: LEDProfile;
	    anti_ice?: LEDProfile;
	    eng_starter?: LEDProfile;
	    apu?: LEDProfile;
	    vacuum?: LEDProfile;
	    hydro_low_pressure?: LEDProfile;
	    aux_fuel_pump?: LEDProfile;
	    parking_brake?: LEDProfile;
	    volt_low?: LEDProfile;
	    doors?: LEDProfile;
	
	    static createFrom(source: any = {}) {
	        return new Leds(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hdg = this.convertValues(source["hdg"], LEDProfile);
	        this.nav = this.convertValues(source["nav"], LEDProfile);
	        this.alt = this.convertValues(source["alt"], LEDProfile);
	        this.apr = this.convertValues(source["apr"], LEDProfile);
	        this.vs = this.convertValues(source["vs"], LEDProfile);
	        this.ap = this.convertValues(source["ap"], LEDProfile);
	        this.ias = this.convertValues(source["ias"], LEDProfile);
	        this.rev = this.convertValues(source["rev"], LEDProfile);
	        this.gear = this.convertValues(source["gear"], LEDProfile);
	        this.master_warn = this.convertValues(source["master_warn"], LEDProfile);
	        this.master_caution = this.convertValues(source["master_caution"], LEDProfile);
	        this.fire = this.convertValues(source["fire"], LEDProfile);
	        this.oil_low_pressure = this.convertValues(source["oil_low_pressure"], LEDProfile);
	        this.fuel_low_pressure = this.convertValues(source["fuel_low_pressure"], LEDProfile);
	        this.anti_ice = this.convertValues(source["anti_ice"], LEDProfile);
	        this.eng_starter = this.convertValues(source["eng_starter"], LEDProfile);
	        this.apu = this.convertValues(source["apu"], LEDProfile);
	        this.vacuum = this.convertValues(source["vacuum"], LEDProfile);
	        this.hydro_low_pressure = this.convertValues(source["hydro_low_pressure"], LEDProfile);
	        this.aux_fuel_pump = this.convertValues(source["aux_fuel_pump"], LEDProfile);
	        this.parking_brake = this.convertValues(source["parking_brake"], LEDProfile);
	        this.volt_low = this.convertValues(source["volt_low"], LEDProfile);
	        this.doors = this.convertValues(source["doors"], LEDProfile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Metadata {
	    name?: string;
	    description?: string;
	
	    static createFrom(source: any = {}) {
	        return new Metadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class Profile {
	    metadata?: Metadata;
	    knobs?: Knobs;
	    leds?: Leds;
	    data?: Data;
	    conditions?: Conditions;
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.metadata = this.convertValues(source["metadata"], Metadata);
	        this.knobs = this.convertValues(source["knobs"], Knobs);
	        this.leds = this.convertValues(source["leds"], Leds);
	        this.data = this.convertValues(source["data"], Data);
	        this.conditions = this.convertValues(source["conditions"], Conditions);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

