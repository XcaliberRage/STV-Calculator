"""
File manipulation functions
"""
import csv
from stvClasses import Candidate, Predictions, Ballot

# Load up a csv file, save three values of information (one int and 2 dictionaries of classes)
# In future would replace predictions with ballots
# and record a dicitonary of classes tracking each VALID voting ballot (provided in the CSV)
def loadCSV(target):

    candidates = {}
    predictions = {}
    parties = {}

    with open(target, newline='') as csvfile:
        linereader = csv.reader(csvfile, delimiter=',')

        # Line 1 contains the seats value
        seats = int(next(linereader)[1])

        # Skip the next two header lines
        skipLines(linereader, 2)

        line = next(linereader)
        # Now we create a dictionary of Candidates starting at line 3
        while line[0] != 'Preferences:':
            # Use the Candidates name as a key.
            # The CSV orders these lines as ID, Candidate, Party, Votes
            # So I have a class that is expecting that order of variables
            # Also the csv file comes with 7 elements when I need 5, so I just lop off the last two
            candidates[line[1]] = Candidate(*line[:4])
            line = next(linereader)

        # After Preferences we expect another line of headers so skip that
        skipLines(linereader, 1)

        # --- CHANGE THIS TO GET BALLOTS INSTEAD ---

        # Generate the dicitonary of predictions
        # Item 0 and 1 are Id and ratio respectively
        # Third is our variable list of parties in order,
        # we also don't need to store more than one '-'
        for line in linereader:
            predictions[line[0]] = Predictions(line[0], line[1],
                                              [item for item in line[2:line.index('-') + 1]])

        # predictions = ballots
        return seats, candidates, predictions


# Helper function for skipping over the CSV lines                                            
def skipLines(fileLine, count):
    for i in range(count):
        next(fileLine)
