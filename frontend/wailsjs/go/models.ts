export namespace main {
	
	export class Match {
	    date: string;
	    turns: number;
	    untrackedWins: number;
	    untrackedLoses: number;
	    won: boolean;
	    streak: number;
	    trophiesFrom: number;
	    trophiesTo: number;
	
	    static createFrom(source: any = {}) {
	        return new Match(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.turns = source["turns"];
	        this.untrackedWins = source["untrackedWins"];
	        this.untrackedLoses = source["untrackedLoses"];
	        this.won = source["won"];
	        this.streak = source["streak"];
	        this.trophiesFrom = source["trophiesFrom"];
	        this.trophiesTo = source["trophiesTo"];
	    }
	}
	export class Profile {
	    isDarkMode: boolean;
	    isLeagueThemed: boolean;
	    rankedPlayed: number;
	    rankedWon: number;
	    matches: Match[];
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isDarkMode = source["isDarkMode"];
	        this.isLeagueThemed = source["isLeagueThemed"];
	        this.rankedPlayed = source["rankedPlayed"];
	        this.rankedWon = source["rankedWon"];
	        this.matches = this.convertValues(source["matches"], Match);
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

