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
  }

  // Get arguments from command line
  population, err := strconv.Atoi(os.Args[1])
  days, err := strconv.Atoi(os.Args[2])
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

  // Simulate one day at a time
  for i := 0; i < days; i++ {
    for j := 0; j < rate; j++ {
      total++
      seed := rand.NewSource(time.Now().UnixNano())
      ran := rand.New(seed)
      num := ran.Intn(100)
      voteChance := ran.Intn(100)
      voteSide := rand.Intn(100)
      // Check what party they are in
      if num < partyLine {  //Democrat
        if voteChance < demEncounter.VoteChance {
          if voteSide < demEncounter.ChanceRepublican {
            repVotes++
          } else {
            demVotes++
          }
        }
      } else {          //Republican
        if voteChance < repEncounter.VoteChance {
          if voteSide < repEncounter.ChanceRepublican {
            repVotes++
          } else {
            demVotes++
          }
        }
      }
    }
  } // End Days loop


  // Calculate how the rest of the population votes(or doesn't)
  for k := total; k < population; k++ {
    total++
    seed := rand.NewSource(time.Now().UnixNano())
    ran := rand.New(seed)
    party := ran.Intn(100)
    voteChance := ran.Intn(100)
    voteSide := rand.Intn(100)

    // Check what party they are in
    if party < partyLine {  // Democrat
      if voteChance < democrat.VoteChance {
        if voteSide < democrat.ChanceRepublican {
          repVotes++
        } else {
          demVotes++
        }
      }
    } else {    // Republican
      if voteChance < republican.VoteChance {
        if voteSide < republican.ChanceRepublican {
          repVotes++
        } else {
          demVotes++
        }
      }
    }
  }


  fmt.Println("Total: ", total)
  fmt.Println("Total Votes: ", demVotes + repVotes)
  fmt.Println("No Votes: ", total-(demVotes+repVotes))
  fmt.Println("Vote Dem: ", demVotes)
  fmt.Println("Vote Rep: ", repVotes)


}
