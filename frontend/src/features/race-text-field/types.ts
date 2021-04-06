type status = 'idle' | 'loading' | 'succeeded' | 'failed';
export type language = 'plain_text' | 'c_cpp' | 'golang' | 'javascript' | 'python';

interface RaceSnippet {
	id: string;
	raceContent: string;
	tokenCount: number;
	language: language;
}

export interface NewPracticeRace {
	snippet: RaceSnippet;
	timeLimit: number;
}

export interface RaceTextFieldState {
	practiceRace: NewPracticeRace;
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
