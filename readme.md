# Movie Recommendation System

This project is a simple movie recommendation system based on the analysis of user preferences and movie characteristics.

## Description

The system provides movie recommendations for users based on their previous interactions with content. For each user, the movies they have watched are analyzed, and a personalized list of recommendations is generated based on this information.

## Installation

### Clone the repository:
```cmd
   git clone https://github.com/Mirzoev-Parviz/movie-recommendation.git
```
### Navigate to the project directory:
```cmd
   cd recommendation-system
```
### Run the application:
```cmd
    go run ./cmd
```


## Usage

The system provides an HTTP API for obtaining recommendations for a specific user.

### Request

```json
{
"user_id": 123
}
```


### Response

```json
[
  {
    "item_id": 456,
    "content_type": "movie",
    "title": "Movie 1",
    "title_orig": "Movie 1",
    "release_year": 2010,
    "genres": ["Action", "Adventure"],
    "countries": ["USA"],
    "for_kids": false,
    "age_rating": 16,
    "studios": ["Studio A"],
    "directors": ["Director X"],
    "actors": ["Actor A", "Actor B"],
    "description": "Description of Movie 1",
    "keywords": ["keyword 1", "keyword 2"]
  },
  {
    "item_id": 789,
    "content_type": "series",
    "title": "Series 1",
    "title_orig": "Series 1",
    "release_year": 2015,
    "genres": ["Drama", "Romance"],
    "countries": ["UK"],
    "for_kids": false,
    "age_rating": 18,
    "studios": ["Studio B"],
    "directors": ["Director Y"],
    "actors": ["Actor C", "Actor D"],
    "description": "Description of Series 1",
    "keywords": ["keyword 3", "keyword 4"]
  }
]
```

## Features
* The system supports different types of content: movies, series, etc.
* Provides personalized recommendations based on user preferences.
* Analyzes movie content and user preferences to generate recommendations.

## Contribution
Any ideas and suggestions are welcome! If you'd like to contribute to the project, please create a pull request or get in touch with us.