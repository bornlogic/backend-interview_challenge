# Interview-Code-Challenge

{{bornlogic tech briefing}}

## Who we want on our Team

We love innovators!
We are looking for entry level engineers to learn and practice software development.
We work hard, play hard and thrive at bornlogic, if you belong to startup way of life, for sure, this is a perfect opportunity!
We want to work with people that build such cool products that they will be proud to show them to their friends and family!


## What should you expect in this code challenge?

This challenges starts your job application at bornlogic.

After you are done, we will review your code, give you feedback and proceed to the next steps of your job application.

## tips
It’s natural to feel stressed during a test, so don’t panic. We wrote down some 'old but gold' tips that may calm you down.

- [ ] Relax yourself mentally and physically
- [ ] Consider alternative solutions
- [ ] If you are stuck, start with the brute force solution
- [ ] Plan out the full solution before you code
- [ ] Keep the big picture in mind

# Let's get started 

Let's pretend we launched the new "Social Media for the Illuminati Secret Society". 
Your first objective as a backend developer is to help us grow our socialmedia API stack, by building new functionalities.


> Keep Calm, do as much as you can, but notice you will have up to 7 days to complete the challenge.


## Developmenty Requirements

### Which technologies can I use?
- C#, Go, Kotlin
- Beginners can do it using Python
- NoSQL (MongoDB, DynamoDB, Elasticsearch)
- SQL (Postgres, CockroachDB)
- Graph (Neo4J, DGraph)

### Should my code be automated tested?
- Yes.

### Step One 

Build a function to validate if a square matrix (NxN) is a 'triangular matrix'. 
{{image}}

## Step Two

Provide a way to access this function publicly. It should be a Web API and a command line interface (CLI).

## Step Three

check this {{swagger link}}

Start the product development following those requirements:

- The following fields from the request must be stored locally: 'Name', 'Date of birth' and 'Notes'.
- Every person also has a counter of how many times they have displayed in the social platform from external API.
- Only Illuminati can access those API, thus it's necessary to send a matrix of numbers (member signature) in a header called ´itanimulli´. If authentication fails, a 404 status code must be returned (not a 403 as correct approach), keeping it secret from eavesdroppers.
- The API must be REST.

It would be great to see these methods, but we would love to hear your proposed list:
- Invite a new illuminati member
- Listing all members (try using a cursor based pagination)
- Search member by name or ID
- Report an outsider (update a flag if the user account doesn't belongs to a real illuminati person)
- Ranking of "displayed member" by month (top five)


## BONUS

- Use best practices notes in 12-factors manifest (https://12factor.net)
- Deploy and share it on a cloud service: Azure, Amazon Web Services, or Google Cloud Plataform 
- Create a docker image
- Use of Behavior Driven Development (BDD) and/or Test Driven Development (TDD)

Δ
