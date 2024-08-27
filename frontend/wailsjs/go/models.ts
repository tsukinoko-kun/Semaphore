export namespace mail {
	
	export class Conversation {
	    account: string;
	    subject: string;
	    participants: string[];
	    lastMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new Conversation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.account = source["account"];
	        this.subject = source["subject"];
	        this.participants = source["participants"];
	        this.lastMessage = source["lastMessage"];
	    }
	}

}

