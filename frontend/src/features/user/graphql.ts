// GraphQL query for getting a NewPracticeRace
export const fetchUserDataGQL = `
	query fetchUserData($userid: String) {
		user(userId: $userid) {
			preferences {
				displayName
			}
			profile {
				averageTPMAllTime
				averageTPMLast10
				maxTPM
				racesCompleted
				racesWon
				totalWordsTyped
			}
		}
	}
`;
