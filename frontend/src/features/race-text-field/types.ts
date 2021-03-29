interface RaceSnippet {
	id: string,
	raceContent: string,
	tokenCount: number
}

export interface NewPracticeRace {
	snippet: RaceSnippet,
	timeLimit: number
}

export interface NewPracticeRaceGQLResponse { 
	data: {
		practiceRace: NewPracticeRace
	}
}

type status = 'idle' | 'loading' | 'succeeded' | 'failed';

export interface RaceTextFieldState {
	practiceRace: NewPracticeRace;
	status: status
}

