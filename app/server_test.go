package poker_test

import (
	poker "github.com/edreg/awesome/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	scores := map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}
	store := poker.NewStubPlayerStore(scores, nil, nil)

	server := poker.NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

//curl -X POST http://localhost:5000/players/Pepper
func TestStoreWins(t *testing.T) {
	store := poker.NewStubPlayerStore(map[string]int{}, nil, nil)

	server := poker.NewPlayerServer(store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := poker.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)
		winCalls := store.GetWinCalls()

		if len(winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(winCalls), 1)
		}

		if winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := poker.NewStubPlayerStore(nil, nil, wantedLeague)

		server := poker.NewPlayerServer(store)

		request := poker.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.GetLeagueFromResponse(t, response.Body)
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(response, t, poker.JsonContentType)
	})
}
