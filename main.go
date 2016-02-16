package main

import (
  "fmt"
  "os"
  "math/rand"
  "time"
  "strconv"
)

type VoteStats struct {
  VoteChance       int
  ChanceRepublican int
  ChanceDemocrat   int
}

func main() {
  // Check to ensure the right number of arguments
  if len(os.Args) != 4 {
    fmt.Println("Usage: vote-sim population days encounter-per-day")
    os.Exit(1)
  }

  // Get arguments from command line
  population, err := strconv.Atoi(os.Args[1])
  // Check if error occured
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  days, err := strconv.Atoi(os.Args[2])
  // Check if error occured
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  rate, err := strconv.Atoi(os.Args[3])
  // Check if error occured
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Declare voting stat blocks
  republican := VoteStats {50, 80, 20}
  repEncounter := VoteStats {60, 60, 40}
  democrat := VoteStats {40, 20, 80}
  demEncounter := VoteStats {70, 10, 90}
  partyLine := 32

  // Variables to hold number of votes
  var total int
  var demVotes int
  var repVotes int

  // Seed random number with time
  seed := rand.NewSource(time.Now().UnixNano())
  ran := rand.New(seed)

  // Simulate one day at a time
  for i := 0; i < days; i++ {
    for j := 0; j < rate; j++ {
      total++
      party := ran.Intn(100)
      voteChance := ran.Intn(100)
      voteSide := rand.Intn(100)
      // Check what party they are in
      if party < partyLine {  //Democrat
        r, d := calcVote(demEncounter, voteChance, voteSide)
        repVotes += r
        demVotes += d
      } else {    //Republican
        r, d := calcVote(repEncounter, voteChance, voteSide)
        repVotes += r
        demVotes += d
      }
    }
  } // End Days loop

  // Calculate how the rest of the population votes(or doesn't)
  for k := total; k < population; k++ {
    total++
    party := ran.Intn(100)
    voteChance := ran.Intn(100)
    voteSide := rand.Intn(100)

    // Check what party they are in
    if party < partyLine {  // Democrat
      r, d := calcVote(democrat, voteChance, voteSide)
      repVotes += r
      demVotes += d
    } else {    // Republican
      r, d := calcVote(republican, voteChance, voteSide)
      repVotes += r
      demVotes += d
    }
  }

  fmt.Println("Total: ", total)
  fmt.Println("Total Votes: ", demVotes + repVotes)
  fmt.Println("No Votes: ", total-(demVotes+repVotes))
  fmt.Println("Vote Dem: ", demVotes)
  fmt.Println("Vote Rep: ", repVotes)
  fmt.Println("Difference", demVotes-repVotes)
  if demVotes - repVotes > 0 {
    fmt.Println("WIN!")
  } else {
    fmt.Println("loss")
  }
}

func calcVote(vote VoteStats, chance, side int) (rep, dem int) {
  if chance < vote.VoteChance {
    if side < vote.ChanceRepublican {
      rep++
    } else {
      dem++
    }
  }
  return
}
