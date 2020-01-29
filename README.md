![alt text](http://branding.bornlogic.com.s3-website-sa-east-1.amazonaws.com/static/files-download/logo/logo-roxo.svg 'Bornlogic')

# Interview backend code challenge

**bornlogic** is a technology startup focused on building digital marketing products. We build a platform to orchestrate thousands of teams creating content everyday. Also, our technology stack handles millions of transactions per day!

We work with people that are passionated about technology, forming a great team of engineers and valuable customers.

## Who we want on our Team

We work hard, play hard and thrive at bornlogic! if you belong to startup way of life, for sure this is a perfect opportunity!

We welcome highly skilled people who want to build amazing products and will be proud to show them to their friends and family! We are also looking for entry-level engineers who want to learn and practice software development using the best practices used by many players in the marketfield.

## What should you expect in this code challenge?

This challenge is the begninning of your job application at bornlogic. 

You have to create a new issue requesting your token access.

After you are done, we will review your code, give you feedback and proceed to the next steps of your job application.


## tips
It’s natural to feel stressed during a test, so don’t panic. We wrote down some good old tips that may calm you down.

- [ ] Relax yourself mentally and physically
- [ ] Consider alternative solutions
- [ ] If you are stuck, start with the brute force solution
- [ ] Plan out the full solution before you code
- [ ] Keep the big picture in mind

# How can I apply for this oportunity?
Fork this repository and pull request your solution. 
You can open an issue at any time if you have any questions.

# Let's get started 

Let's pretend we launched the new "Social Media for the Illuminati Secret Society". 
Your first objective as a backend developer is to help us grow our socialmedia API stack by building new functionalities.

> Keep Calm, do as much as you can, but keep in mind that you have up to 7 days to complete the challenge.

## Development Requirements

### Which technologies can I use?
- C#, Go, Kotlin
- Beginners can use Python
- NoSQL (MongoDB, DynamoDB, Elasticsearch)
- SQL (Postgres, CockroachDB)
- Graph (Neo4J, DGraph)

### Should my code be automated tested?
- Yes.

### Step One 

Build a function to validate if a square matrix (NxN) is a 'triangular matrix'. 

![alt This matrix is upper triangular](https://wikimedia.org/api/rest_v1/media/math/render/svg/d2ed83943adf49954804bdfa12ca1bb4b278a64d)
This matrix is upper triangular

![alt This matrix is low triangular](https://wikimedia.org/api/rest_v1/media/math/render/svg/3fc7115ee860d2758e890e4de5217aafb89d90e6)
This matrix is lower triangular.

## Step Two

Provide a way to access this function publicly. It should have both a Web REST API and a command line interface (CLI).

## Step Three

It's show time! You should access this API:

https://us-east4-sandbox-254919.cloudfunctions.net/interview-code-challenge-api

Response example:

     {
        "people": [
            {
                "summary_metrics": {
                    "posts": [
                        {
                            "photo": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 495,
                                    "likes": 5466,
                                    "shares": 1445
                                }
                            },
                            "video": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 1528,
                                    "completed_view": 9106,
                                    "likes": 6258
                                }
                            }
                        }
                    ]
                },
                "Name": "Jhon Doe",
                "DateOfBirth": "1985-02-02",
                "Notes": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque arcu tortor, ornare eget ..."
            },
            {
                "summary_metrics": {
                    "posts": [
                        {
                            "photo": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 8047,
                                    "likes": 9947,
                                    "shares": 1445
                                }
                            },
                            "video": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 8287,
                                    "completed_view": 9106,
                                    "likes": 2888
                                }
                            }
                        }
                    ]
                },
                "Name": "Franky Kelley",
                "DateOfBirth": "1976-10-05",
                "Notes": "Nulla quis purus odio. Suspendisse aliquet vehicula leo, sed dignissim dolor pellentesque ..."
            },
            {
                "summary_metrics": {
                    "posts": [
                        {
                            "photo": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 2790,
                                    "likes": 3015,
                                    "shares": 1445
                                }
                            },
                            "video": {
                                "URL": "http://medias-challenge.com/media.jpeg",
                                "Metrics": {
                                    "clicks": 5541,
                                    "completed_view": 9106,
                                    "likes": 408
                                }
                            }
                        }
                    ]
                },
                "Name": "Lacie Scott",
                "DateOfBirth": "1999-07-12",
                "Notes": "Integer nec tellus suscipit, vestibulum urna at, iaculis nisl. Maecenas ante est..."
            }
        ]
    }


Now you can start the product development following these requirements:

- The following fields from the request must be stored locally: 'Name', 'Date of birth' and 'Notes'.
- Every person also has a counter of how many times they have displayed in the social platform from external API.
- Only Illuminati can access those API, so it's necessary to send a matrix of numbers (member signature) in a header called ´itanimulli´.

It would be great to see these methods, but we would love to hear your proposed list:
- Invite a new illuminati member
- Listing all members (try using a cursor based pagination)
- Search member by name or ID
- Report an outsider (update a flag if the user account doesn't belongs to a real illuminati person)
- Ranking of "displayed member" by month (top five)


## BONUS

- Use best practices notes in [12-factors manifest](https://12factor.net) 
- Deploy and share it on a cloud service: Azure, Amazon Web Services, or Google Cloud Plataform 
- Create a docker image
- Use of Behavior Driven Development (BDD) and/or Test Driven Development (TDD)

Δ
