package domain

// GameFeedDTO represents any GameFeed message with its payload and metadata
type GameFeedDTO struct {
	Data     GameData `json:"data"`
	Finished bool     `json:"finished"`
}

// GameData represents a feed line of a tennis game
type GameData struct {
	TourneyID        string `json:"tourney_id"`
	TourneyName      string `json:"tourney_name"`
	TourneyDate      string `json:"tourney_date"`
	Surface          string `json:"surface"`
	WinnerID         string `json:"winner_id"`
	LoserID          string `json:"loser_id"`
	Score            string `json:"score"`
	BestOf           string `json:"best_of"`
	Round            string `json:"round"`
	Minutes          string `json:"minutes"`
	WAce             string `json:"w_ace"`
	WDf              string `json:"w_df"`
	WSvpt            string `json:"w_svpt"`
	W1stIn           string `json:"w_1st_in"`
	W1stWon          string `json:"w_1st_won"`
	W2ndWon          string `json:"w_2nd_won"`
	WSvGms           string `json:"w_sv_gms"`
	WbpSaved         string `json:"w_bp_saved"`
	WbpFaced         string `json:"w_bp_faced"`
	Lace             string `json:"l_ace"`
	Ldf              string `json:"l_df"`
	Lsvpt            string `json:"l_svpt"`
	L1stIn           string `json:"l_1st_in"`
	L1stWon          string `json:"l_1st_won"`
	L2ndWon          string `json:"l_2nd_won"`
	LSvGms           string `json:"l_sv_gms"`
	LbpSaved         string `json:"l_bp_saved"`
	LbpFaced         string `json:"l_bp_faced"`
	WinnerRank       string `json:"winner_rank"`
	WinnerRankPoints string `json:"winner_rank_points"`
	LoserRank        string `json:"loser_rank"`
	LoserRankPoints  string `json:"loser_rank_points"`
}

// BuildGameFeedDTO builds a GameFeedDTO from a slice
func BuildGameFeedDTO(record []string) GameFeedDTO {
	return GameFeedDTO{
		Data: GameData{
			TourneyID:        record[0],
			TourneyName:      record[1],
			TourneyDate:      record[2],
			Surface:          record[3],
			WinnerID:         record[4],
			LoserID:          record[5],
			Score:            record[6],
			BestOf:           record[7],
			Round:            record[8],
			Minutes:          record[9],
			WAce:             record[10],
			WDf:              record[11],
			WSvpt:            record[12],
			W1stIn:           record[13],
			W1stWon:          record[14],
			W2ndWon:          record[15],
			WSvGms:           record[16],
			WbpSaved:         record[17],
			WbpFaced:         record[18],
			Lace:             record[19],
			Ldf:              record[20],
			Lsvpt:            record[21],
			L1stIn:           record[22],
			L1stWon:          record[23],
			L2ndWon:          record[24],
			LSvGms:           record[25],
			LbpSaved:         record[26],
			LbpFaced:         record[27],
			WinnerRank:       record[28],
			WinnerRankPoints: record[29],
			LoserRank:        record[30],
			LoserRankPoints:  record[31],
		},
		Finished: false,
	}
}

// BuildEndMessage returns a message representing the end of the feed
func BuildEndMessage() GameFeedDTO {
	return GameFeedDTO{
		Data:     GameData{},
		Finished: true,
	}
}
