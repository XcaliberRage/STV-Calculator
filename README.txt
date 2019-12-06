This program is a method to handle STV voting

stv.py expects to find a series of .csv files
Each of which will be specifically formatted to present voting data for a Super constituence
This includes:*
Number of intended seats
A Table of Candidates with their party memebership and unique ID (counting integer)
A Table of all valid ballots with unique ID and selected candidates in order of preference, '-' indicates no further selections made

The software will then print out round by round the outcomes of the STV rounds until all seats have a winner allocated

*
-------- IN REALITY ---------
Instead of Ballots, the program uses "predictions by party"
And the candidates have a record of their original choice
This is because the program currently simulates STV using FPTP voting information
Therefore a series of predictions on voter preferences were made by party (indluing ratio splits)
This is used to then fabricate ballots that would always put members of one party before selecting another
The ballots are then tracked in lumps using their ratio against their first choice MP's original votes


TODO: Save a new .csv file reporting the outcome