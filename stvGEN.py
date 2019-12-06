"""
These functions are concerned with generating Ballots based off of the CSV inputs
and then used to track and count votes.
In reality you'd be loading CSV files that contain each individual vote rather
than generating ballots based on predictions.
Therefore, this entire file would then be deleted and Make_Ballots would be
a part of the loadCSV function
"""

from stvClasses import Candidate, Predictions, Ballot

# Iterate over the candidates, making a Ballot for each combination of prediction that starts with this candidate's party
# The ballot will have a proportion of this candidates votes and a complete list of ordered preference of each MP in a party they selected
# Seocndary to party order is MP order which is by number of initial votes
def Make_Ballots(candidates, predictions):

    # I want to group the candidates by party and order them from highest votes to lowest
    # I only do this to help with artificial ballot generation, normally this would not be necessary at all
    parties = get_Parties(candidates)

    for party in parties:
        print(party)
        print(parties[party])
        print()


    # So iterate over each candidate
    ballots = {}
    ballotID = 0

    for candidate in candidates:
        # Note this candidate's party
        party = candidates[candidate].party

        # Iterate over the predictions
        for prediction in predictions:
            # But only the ones starting with this party
            if predictions[prediction].prefs[0] == party:
                # Do we create a new ballot
                order = list([candidate])

                # but first we need to create a list that order's this user's preferences
                for pGroup in predictions[prediction].prefs:
                    # At end of list we just dump this into the end of prefs (we use it to trigger Ballot deletion)
                    if pGroup == 'ABSTAIN':
                        order.append('ABSTAIN')
                    else:
                        for candy in parties[pGroup]:
                            # We avoid duplicates
                            if candy != candidate:
                                order.append(candy)
                                
                ballots[ballotID] = Ballot(ballotID,
                                           (candidates[candidate].votes * predictions[prediction].ratio),
                                           order)
                # Then just increment the ballotID because we will need a new one
                ballotID += 1

    return ballots


# Using the list of candidates, group them into lists of same party, ordered by highest votes to lowest
def get_Parties(candidates):

    parties = {}
    
    candVotes = {candidate: candidates[candidate].votes for candidate in candidates}

    # In order, input the candidates from highest votes to lowest
    # Iterate over the dictionary
    for i in range(len(candVotes)):
        # Identify the most popular candidate
        target = max(candVotes, key=candVotes.get)

        # When the party is new, you have to create it, i.e. if you try to directly append
        # to the list in parties[party] you'll just get a KeyError:
        # So we'll first check if it exists and if not create a 1 item list and plop it in
        party = candidates[target].party
        if party in parties:
            parties[party].append(target)
        else:
            parties[party] = list([target])

        # Delete this candidate from candVotes
        del candVotes[target]

    return parties
