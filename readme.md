## Shitpost

REST API to create shitposts. Purely made for learning purposes.

The goal here is to create a REST API that allows users to create shitposts. The shitposts will be stored in a database and can be retrieved, created and deleted. The API will be built using Golang and the database will be SQLite.

We'll try to follow the best practices for building REST APIs and we'll also try to make the code as clean as possible. We'll avoid using any external libraries we'll prefer the standard library instead.

### Requirements

Features:
    - Shitposts are anonymous and can be created by anyone.
    - To delete a shitpost, the user must provide a passcode that was generated when the shitpost was created.

API Design:
    - List API must support pagination.
    - Shitpost content must be limited to 500 characters.


### Usage

#### Running the server

```bash
./shitpost server
```

#### Running migrations

```bash
./shitpost migrations up
./shitposts migrations down
./shitposts migrations to 0001
./shitposts migrations drop
```

### Endpoints

#### Create a Shitpost

```
POST /shitposts

{
    "title": "Omae wa mou shindeiru",
    "content": "What the fuck did you say to me you motherfucker? I'll let you know that I've graduated on the top of my class in the Indian Naval Forces. I've been involved in countless secret raids against Pakistan and I've sent over 300+ cunts to their death. I'm experienced in guerrilla warfare and I'm the elite sniper of the Indian Army. You are just another target to me. I'll not only wipe your existence from the world, but the entire universe will hear you getting wiped out. Listen, motherfucker. What do you think, you can just tell me whatever over the internet and I'll keep listening? Think again, cunt. As I'm writing this officers at RAW are already tracking your IP address. Be prepared, kiddo. Because the darkness that's coming on you, the likes of which no one has seen in history. Such a darkness that will finish the thing you call your "life". You're done. I can be anywhere, anytime and finish you in over 700 ways just with my bare hands. Not only I'm trained in armed combat, but I can use the entire arensal of the Indian Army, which I will use to the fullest extent to wipe your existence which will be remembered across the world, greater than the distance of the seven seas. Maybe you should have thought before posting your "clever" comment, which would've had such grave consequences. But it's too late now. You're dead motherfucker."
}
```

#### Get a Shitpost

```
GET /shitposts/{id}
```

### List shitposts

```
GET /shitposts
```

```
GET /shitposts?page=2
```

#### Delete a Shitpost

```
DELETE /shitposts/{id}

{
    "secret": "my-secret"
}
```

#### OpenAPI docs

```
GET /swagger/
```
