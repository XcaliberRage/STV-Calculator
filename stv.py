"""
Take a series of political candidates and their votes.
Use a prediction of voter preferences to simulate an STV for X seats
In future, change handling of predicitons to instead reading real votes
"""
import os, math
from stvClasses import Candidate, Predictions, Ballot
from stvCSV import loadCSV
from stvGEN import Make_Ballots, get_Parties

# Start the system by analysing the dir and pulling each CSV file (containing the STV data)
# Each file should be one Super Constituency of data in a standardised format
def main():

    # Get the file names of each CSV file to be worked with
    directory = os.fsencode(os.curdir)
    csvFiles = []

    for file in os.listdir(directory):
        filename = os.fsdecode(file)
        if filename.endswith(".csv"):
            csvFiles.append(filename)
            continue
        else:
            continue

    # For each csv file, generate an output of results
    for superConstituency in csvFiles:
        Simulate_STV(superConstituency)
            

# The bulk operation, on each file we load the info, run the simulations and then save a csv file with the info
def Simulate_STV(superConstituency):

    name = superConstituency.replace('.csv', '')
    seats = 0
    candidates = {}
    predictions = {}

    # First pull the data from the CSV file
    # seats is an integer
    # candidates is a dictionary of Candidate classes
    # predictions is a dictionary of Prediction classes <--- Would be deleted for real STV
    #
    # For real STV: ballots would be given here
    #
    seats, candidates, predictions = loadCSV(superConstituency)

    print(name, ' ---')
    print('Seats: ', seats)
    print('Candidates:')
    for i in candidates:
        candidates[i].report()

    print()
    print('Voter preferences ----')
    for i in predictions:
        predictions[i].report()
    
    # I should have suitable data to create a list of Ballots using a combination of the estimated preferences
    # And original vote data
    # This function is in place of a function that loadCSV would use in order to report the actual ballots
    ballots = Make_Ballots(candidates, predictions)
    balLen = len(ballots)

    # Create a modifiable list of candidates using their name as a key and their originaly votes as their votes
    candidatesLive = {candidate: 0 for candidate in candidates}
    for cand in candidatesLive:
        print(cand, candidatesLive[cand])
    print()

    # Election_Rounds will return a tuple:
    # First list of dictionaries indicating the seat winners, their votes and what round they won
    # Second a dictionary of eliminated MPs and what round they were eliminated
    winners, eliminated = Election_Rounds(candidatesLive, ballots, seats)

    # Yay we've (in theorey, run a round of STV)
    # Now all we need to do is report the data
    print('----------------------')
    print('Winners:')
    for i in range(len(winners)):
        winners[i]['party'] = candidates[winners[i]['name']].party
        print('Seat ', i+1, ': ', winners[i]['name'],
              ' for ', winners[i]['party'],
              ' with ', '{:.2f}'.format(winners[i]['votes']),
              ' votes (R', winners[i]['round'], ')')


# This function initiates and executes the STV rounds
def Election_Rounds(candidatesLive, ballots, seats):

    winners = []
    eliminated = {}
    oldScore = {}
    roundCount = 0

    # If we've filled all seats we're done, same if there's no more candidates standing
    while True:
        winLen = len(winners)
        canLen = len(candidatesLive)

        if winLen >= seats or canLen == 0:
            print('ALL SEATS FILLED, COUNTS COMPLETE')
            print()
            break
        # New round
        totalVotes = 0
        numBal = len(ballots)
        totalVotes = tallyVotes(ballots)
        roundCount += 1

        print('-----------')
        print('Round: ', roundCount)
        print('Total votes = ', '{:.2f}'.format(totalVotes))
        print()

        # If there's only one candidate left and we've got here (because there's empty seats)
        # They will get the last seat awarded
        if len(candidatesLive) == 1:
            for cand in candidatesLive:
                name = cand
                votes = candidatesLive[cand]
                
            winners.append({
                'name': name,
                'round': roundCount,
                'votes': int(votes)
                })
            print('Last candidate standing for empty seat = auto win')
            
        else:
            # For each candidate, tally their live votes
            for candidate in candidatesLive:
                # Score the candidates
                oldScore[candidate] = candidatesLive[candidate]
                candidatesLive[candidate] = 0
                for ballot in ballots:
                    # CHANGE ballots[ballot].wtdVotes to ballots[ballot].weight FOR REAL STV
                    candidatesLive[candidate] += ballots[ballot].wtdVotes if ballots[ballot].pick[0] == candidate else 0

            print('Standings:')
            print('Current winners: ', ', '.join([winner['name'] for winner in winners]))
            print()
            for candidate in candidatesLive:
                print(candidate, ': ', int(candidatesLive[candidate]), end=' ')
                gained = (int(candidatesLive[candidate]) - int(oldScore[candidate]))
                if gained > 0:
                    print('(gaining ', gained, ' votes)')
                else:
                    print()
            print()
                    
            # Get the droop Quota
            dQ = droop(totalVotes, seats)
            print('Total Votes = ','{:.2f}'.format(totalVotes))
            print('Droop Quota = ', dQ)
            print()

            # Look for candidates with votes >= dQ
            if doQualify(candidatesLive, dQ):
                # Identify winners and award seats
                # Handle surplus
                print('Candidates qualify:')
                qualifiers = placements(candidatesLive, ballots, dQ)

                # Store the winners from this round into the list of winners with votes to win and the round
                for candidate in qualifiers:
                    if len(winners) < seats:
                        print(candidate, ' with ', dQ, ' votes')
                        winners.append({
                            'name': candidate,
                            'round': roundCount,
                            'votes': int(dQ)})
            else:
                # Eliminate the loser
                eliminated[elimination(candidatesLive, ballots)] = roundCount

            print('ROUND END')
            print()

        
    # Election end
    return winners, eliminated
                

# Identify the candidate with fewest votes and eliminate them returning the candidate's name
def elimination(candidatesLive, ballots):

    # Identify the loser
    loser = min(candidatesLive, key=candidatesLive.get)
    loserVote = candidatesLive[loser]

    # Erase the loser from all existence
    del candidatesLive[loser]
    elims = []
    for ballot in ballots:
        ballots[ballot].ElimCand(loser)

    killAbstains(ballots)

    print('Elimination: ', loser)

    return loser


# Placements looks for the winners, granting seats in order of highest to lowest
# It also triggers and handles weight redistributions
def placements(candidatesLive, ballots, dQ):
    
    quals = []
    while True:
        best = max(candidatesLive, key=candidatesLive.get)
        bestVote = candidatesLive[best]
        if candidatesLive[best] >= dQ:
            # Add the candidate to the list of qualifiers
            quals.append(best)

            # Eliminate the candidate from each ballot
            # But check to see if this ballot was part of the qualifying votes
            # If so change the weight
            for ballot in ballots:
                if ballots[ballot].pick[0] == best:
                    ballots[ballot].ChangeWeight(dQ, candidatesLive[best])

                ballots[ballot].ElimCand(best)
                
            killAbstains(ballots)
            del candidatesLive[best]
        else:
            break

    return quals


# Each time we remove candidates, we also need to indetify if we delete any ballots
# But we can't do it while iterating over ballots so we have to do it after
def killAbstains(ballots):

    elims = []
    for ballot in ballots:
        if ballots[ballot].pick[0] == 'ABSTAIN':
            elims.append(ballot)

    for bal in elims:
        del ballots[bal]





# Helper function to tally up the votes across all ballots after weighting
def tallyVotes(ballots):
    votes = 0
    for ballot in ballots:
        pick = ballots[ballot].pick[0]
        if ballots[ballot].pick[0] != 'ABSTAIN':
            # CHANGE ballots[ballot].wtdVotes TO ballots[ballot].weight FOR REAL STV
            votes += ballots[ballot].wtdVotes

    return votes

# Helper function to get the dQ
def droop(votes, seats):
    return int(math.ceil((votes/(seats+1)+1)))


# Helper function to identify the if there are any qualifiers this round
def doQualify(candidatesLive, dQ):
        
        for candidate in candidatesLive:
            if candidatesLive[candidate] >= dQ:
                return True

        return False
    

if __name__ == '__main__':
    main()
