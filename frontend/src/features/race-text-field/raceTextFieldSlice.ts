import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import axios from 'axios'

// redux prefix for this slice
const RACE_TEXT_FIELD = 'raceTextField';

// GraphQL query for getting a NewPracticeRace
const newPracticeRaceGQL = `
  query getPracticeRace {
    practiceRace {
      snippet {
        id
        raceContent
        tokenCount
      }
      timeLimit
    }
  }
`;

interface RaceSnippet {
	id: string,
	raceContent: string,
	tokenCount: number
}

interface NewPracticeRace {
	snippet: RaceSnippet,
	timeLimit: number
}

interface NewPracticeRaceGQLResponse { 
	data: {
		practiceRace: NewPracticeRace
	}
}

// fetches a new practice race from the backend using graphql
export const fetchNewPracticeRace = createAsyncThunk<NewPracticeRace>(
	`${RACE_TEXT_FIELD}/fetchNewPracticeRace`, 
	async () => {
		const response = await axios.post<NewPracticeRaceGQLResponse>(
			'http://localhost:8080/graphql', 
			{ query: newPracticeRaceGQL }
		);
		return response.data.data.practiceRace;
	}
);

type status = 'idle' | 'loading' | 'succeeded' | 'failed';

interface RaceTextFieldState {
	practiceRace: NewPracticeRace;
	status: status
}

const initialState: RaceTextFieldState = {
	practiceRace: {
		snippet: {
			id: '0',
			raceContent: '',
			tokenCount: 0
		},
		timeLimit: 0
	},
	status: 'idle'
}

const raceTextFieldSlice = createSlice({
	name: 'raceTextField',
	initialState,
	reducers: {},
	extraReducers: builder => {
		builder.addCase(fetchNewPracticeRace.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.practiceRace = action.payload;
		});
		builder.addCase(fetchNewPracticeRace.rejected, (state, action) => {
			console.log(action);
		})
	}
})

export default raceTextFieldSlice.reducer;

export const selectSnippet = (state: RootState) => state.raceTextField.practiceRace.snippet;
