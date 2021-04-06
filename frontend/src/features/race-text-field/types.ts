type status = 'idle' | 'loading' | 'succeeded' | 'failed';
export type language = 'plain_text' | 'c_cpp' | 'golang' | 'javascript' | 'python';

interface RaceSnippet {
	id: string;
	raceContent: string;
	tokenCount: number;
	language: language;
}

export interface Race {
	snippet: RaceSnippet;
	timeLimit: number;
	typedSoFar: string;
}

export interface RaceFieldState {
	race: Race;
	status: status;
}

// types used in the graphql response for getting a new practice race
interface GQLRaceSnippet {
	id: string;
	raceContent: string;
	tokenCount: number;
	language: number;
}

interface NewGQLPracticeRace {
	snippet: GQLRaceSnippet;
	timeLimit: number;
}

export interface NewPracticeRaceGQLResponse { 
	data: {
		practiceRace: NewGQLPracticeRace;
	}
}
