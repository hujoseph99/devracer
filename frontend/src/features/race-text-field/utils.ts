import { language, NewPracticeRace, NewPracticeRaceGQLResponse } from "./types";

// maps an int from the backend to a language
const mapNumberToLanguage = (language: number): language => {
	switch(language) {
		case 1:
			return 'c_cpp';
		case 2:
			return 'golang';
		case 3:
			return 'javascript';
		case 4:
			return 'python';
		default:
			return 'plain_text'
	}
}

// maps the graphql response into our own NewPracticeRace type. The only real change
// that we are making is to map the language to something usable on our end
export const mapGQLPracticeRaceToNewPracticeRace = (res: NewPracticeRaceGQLResponse): NewPracticeRace => {
	let practiceRace = res.data.practiceRace;
	let snippet = practiceRace.snippet
	return {
		snippet: {
			id: snippet.id,
			raceContent: snippet.raceContent,
			tokenCount: snippet.tokenCount,
			language: mapNumberToLanguage(snippet.language)
		},
		timeLimit: practiceRace.timeLimit
	};
}
