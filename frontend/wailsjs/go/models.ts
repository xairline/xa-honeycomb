export namespace pkg {
	
	export class Data {
	    dataref_str?: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dataref_str = source["dataref_str"];
	    }
	}
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
	export class Dataref {
	    dataref_str?: string;
	    operator?: string;
	    threshold?: number;
	    index?: number;
	
	    static createFrom(source: any = {}) {
	        return new Dataref(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dataref_str = source["dataref_str"];
	        this.operator = source["operator"];
	        this.threshold = source["threshold"];
	        this.index = source["index"];
	    }
	}
	export class BravoProfile {
	    profile_type?: string;
	    condition?: string;
	    datarefs?: Dataref[];
	    commands?: Command[];
	    data?: Data[];
	
	    static createFrom(source: any = {}) {
	        return new BravoProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.profile_type = source["profile_type"];
	        this.condition = source["condition"];
	        this.datarefs = this.convertValues(source["datarefs"], Dataref);
	        this.commands = this.convertValues(source["commands"], Command);
	        this.data = this.convertValues(source["data"], Data);
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
	
	
	
	export class Profile {
	    name: string;
	    ap_hdg?: BravoProfile;
	    ap_vs?: BravoProfile;
	    ap_alt?: BravoProfile;
	    ap_ias?: BravoProfile;
	    ap_crs?: BravoProfile;
	    bus_voltage?: BravoProfile;
	    hdg?: BravoProfile;
	    nav?: BravoProfile;
	    alt?: BravoProfile;
	    apr?: BravoProfile;
	    vs?: BravoProfile;
	    ap?: BravoProfile;
	    ias?: BravoProfile;
	    rev?: BravoProfile;
	    ap_state?: BravoProfile;
	    gear?: BravoProfile;
	    retractable_gear?: BravoProfile;
	    master_warn?: BravoProfile;
	    master_caution?: BravoProfile;
	    fire?: BravoProfile;
	    oil_low_pressure?: BravoProfile;
	    fuel_low_pressure?: BravoProfile;
	    anti_ice?: BravoProfile;
	    eng_starter?: BravoProfile;
	    apu?: BravoProfile;
	    vacuum?: BravoProfile;
	    hydro_low_pressure?: BravoProfile;
	    aux_fuel_pump?: BravoProfile;
	    parking_brake?: BravoProfile;
	    volt_low?: BravoProfile;
	    doors?: BravoProfile;
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.ap_hdg = this.convertValues(source["ap_hdg"], BravoProfile);
	        this.ap_vs = this.convertValues(source["ap_vs"], BravoProfile);
	        this.ap_alt = this.convertValues(source["ap_alt"], BravoProfile);
	        this.ap_ias = this.convertValues(source["ap_ias"], BravoProfile);
	        this.ap_crs = this.convertValues(source["ap_crs"], BravoProfile);
	        this.bus_voltage = this.convertValues(source["bus_voltage"], BravoProfile);
	        this.hdg = this.convertValues(source["hdg"], BravoProfile);
	        this.nav = this.convertValues(source["nav"], BravoProfile);
	        this.alt = this.convertValues(source["alt"], BravoProfile);
	        this.apr = this.convertValues(source["apr"], BravoProfile);
	        this.vs = this.convertValues(source["vs"], BravoProfile);
	        this.ap = this.convertValues(source["ap"], BravoProfile);
	        this.ias = this.convertValues(source["ias"], BravoProfile);
	        this.rev = this.convertValues(source["rev"], BravoProfile);
	        this.ap_state = this.convertValues(source["ap_state"], BravoProfile);
	        this.gear = this.convertValues(source["gear"], BravoProfile);
	        this.retractable_gear = this.convertValues(source["retractable_gear"], BravoProfile);
	        this.master_warn = this.convertValues(source["master_warn"], BravoProfile);
	        this.master_caution = this.convertValues(source["master_caution"], BravoProfile);
	        this.fire = this.convertValues(source["fire"], BravoProfile);
	        this.oil_low_pressure = this.convertValues(source["oil_low_pressure"], BravoProfile);
	        this.fuel_low_pressure = this.convertValues(source["fuel_low_pressure"], BravoProfile);
	        this.anti_ice = this.convertValues(source["anti_ice"], BravoProfile);
	        this.eng_starter = this.convertValues(source["eng_starter"], BravoProfile);
	        this.apu = this.convertValues(source["apu"], BravoProfile);
	        this.vacuum = this.convertValues(source["vacuum"], BravoProfile);
	        this.hydro_low_pressure = this.convertValues(source["hydro_low_pressure"], BravoProfile);
	        this.aux_fuel_pump = this.convertValues(source["aux_fuel_pump"], BravoProfile);
	        this.parking_brake = this.convertValues(source["parking_brake"], BravoProfile);
	        this.volt_low = this.convertValues(source["volt_low"], BravoProfile);
	        this.doors = this.convertValues(source["doors"], BravoProfile);
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

