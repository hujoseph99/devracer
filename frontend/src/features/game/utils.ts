import { language, SnippetResponse, SnippetState } from "./types";

const mapLanguage = (language: number): language => {
	switch (language) {
		case 1:
			return 'c_cpp';
		case 2:
			return 'golang';
		case 3:
			return 'javascript';
		case 4:
			return 'python';
		default:
			return 'plain_text';
	}
}

export const transformSnippetResponse = (response: SnippetResponse): SnippetState => {
	return {
		...response,
		language: mapLanguage(response.language)
	}
}
