"""
File containing the 3 classes needed for STV counting
"""


# Class for Political Candidates include unique ID, name,
# affiliated party and their voteshare in float form
class Candidate:
    def __init__(self, _id, name, party, votes):
        self.id = int(_id)
        self.name = name
        self.party = party
        self.votes = float(votes)

    def report(self):
        print('{:02d}'.format(self.id), ' - ', self.name, ' - ',
              self.party, ': ', '{:.2f}'.format(self.votes), 'votes')


# Class for Predicted vote preferences by party and their ratio
class Predictions:
    def __init__(self, _id, ratio, prefs):
        self.id = int(_id)
        self.ratio = float(ratio)
        self.prefs = [i.replace('-', 'ABSTAIN') for i in prefs]

    def report(self):
        print('{:02d}'.format(self.id), ' - ', '{:.02f}'.format(self.ratio * 100), ': ',
              ', '.join(self.prefs))


# Class containing Ballot information, this includes:
# An ID, naturally
# A list of ordered preferences by MP
# A vote count, intialised as a ratio of their first choice MP's original votes
# A vote weight intialised at 1
class Ballot:
    def __init__(self, _id, votes, orderdPref):
        self.id = int(_id)
        self.weight = 1.00
        # 'VOTES' AND 'WTDVOTES' WOULD NOT BE USED WHEN REAL BALLOTS ARE COUNTED
        # Weight would instead be counted as vote (every ballot is one vote) 
        self.votes = float(votes)
        self.wtdVotes = float(votes)
        ####
        
        self.pick = orderdPref

    # Modify these Ballot's weights as per the accepted droop formula
    # Then update the Weighted Votes of these Ballots
    def ChangeWeight(self, dQ, votes):
        self.weight *= ((votes-dQ)/votes)
        # DELETE FOR REAL STV, WEIGHTS WOULD BE NORMALLY USED
        self.wtdVotes = self.votes * self.weight

    # Remove a candidate from the pick
    def ElimCand(self, candidate):
        if candidate in self.pick:
            self.pick.remove(candidate)
            
