# Problem Title:
Character Generator from Trait Collection

# Problem Statement:

You are given a **collection** that contains available character traits. The `collection` is a JSON object structured like this:

```json
{
  "collection": {
    "size": 5,
    "trait": [
      {
        "trait_type": "ear",
        "values": ["small", "big", "pointy", "round"]
      },
      {
        "trait_type": "mouth",
        "values": ["smile", "frown", "neutral", "open"]
      },
      {
        "trait_type": "nose",
        "values": ["small", "big", "sharp", "flat"]
      }
    ]
  }
}
```

Each `trait` has:
- a `trait_type` (e.g., "ear", "mouth", "nose")
- a list of possible `values`

You are tasked with **generating characters**.
Each character must have an **`attributes`** array where:
- Each element of `attributes` is an object with:
    - `trait_type` (string)
    - `value` (string, randomly selected from available values for that trait)

Your goal is to **generate `n` random characters** following the above structure.

# Input:

- A JSON object `collection` (as above).
- An integer `n`, the number of characters to generate.

# Output:

- An array of `n` character objects.
  Each character object must have:
    - `"name"`: (string, e.g., "Character 1", "Character 2", etc.)
    - `"attributes"`: (array of { `trait_type`, `value` } objects)

# Constraints:

- 1 \u2264 n \u2264 1000
- Each trait must be selected independently.
- The value for each trait must be randomly chosen from its available `values`.

# Example:

## Input

```json
{
  "collection": {
    "size": "large",
    "trait": [
      {
        "trait_type": "ear",
        "values": ["small", "big", "pointy", "round"]
      },
      {
        "trait_type": "mouth",
        "values": ["smile", "frown", "neutral", "open"]
      },
      {
        "trait_type": "nose",
        "values": ["small", "big", "sharp", "flat"]
      }
    ]
  }
}
```

n = 2

## Output (One possible output)

```json
[
  {
    "name": "Character 1",
    "attributes": [
      {
        "trait_type": "ear",
        "value": "pointy"
      },
      {
        "trait_type": "mouth",
        "value": "neutral"
      },
      {
        "trait_type": "nose",
        "value": "small"
      }
    ]
  },
  {
    "name": "Character 2",
    "attributes": [
      {
        "trait_type": "ear",
        "value": "big"
      },
      {
        "trait_type": "mouth",
        "value": "smile"
      },
      {
        "trait_type": "nose",
        "value": "sharp"
      }
    ]
  }
]
```

(Note: Since the selection is random, output can vary every time.)

# Follow-Up:

- How would you ensure no duplicate characters are created if required?
- How would you generate weighted random traits if some values are more common than others?

