// GraphQL query for getting a NewPracticeRace
export const newPracticeRaceGQL = `
  query getPracticeRace {
    practiceRace {
      snippet {
        id
        raceContent
        tokenCount
        language
      }
      timeLimit
    }
  }
`;
