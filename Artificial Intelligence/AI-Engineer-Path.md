# AI Engineer Path

- "Artificial Intelligence could be more profound than fire, electricty, or the internet." - Sundar Pichai

<br>

- AI engineer does not refer to a person that train and build the AI models, that is the AI Researcher or Machine Learning Engineer.
- AI engineers use these models to build applications.

<br>
<br>
<hr>

# Intro to AI Engineering
- OpenAI launch 2015, but burst into public around 2022 
- Revolutionized our concept of what AI can do.


### Large Language Model
- is an algorithm that uses  training data to recognize patterns and make predictions or decisions.

<br>

### Models
- model snapshot
- context length
- cut-off date
- memory 


<br>
## Prompt Engineering
- designing inputs for generative AI tools to produce optimal outputs.

<br>
## Tokens
- https://platform.openai.com/tokenizer
- chunk of text of no specific length
- prompt_tokens, completion_tokens, total_tokens
- finish_reason


### Why do tokens matter?
- Tokens cost credit
- The more token you use, the more money you spend.
- Tokens need processing
- The more token you use , the more lag time you get, the slower your app will be.
- Keeping token numbers low saves your users time and saves you money.


### Temperature
- Controls how daring output is
- Can be set to 0 to 2
- Default to 1
- Models are not deterministic
- Lower Temperature
  - Safer, more predictable
  - Good for factual output

- High Temperature
  - More dangerous, less predictable
  - Good for creative output


### Zero shot approach
- Ask what we want without any examples.

### Few shot approach
- Provide the model with one or more examples of  what we want.


- Example
```javascript

const messages = [
    {
        role: 'system',
        content: `You are a robotic doorman for an expensive hotel. When a customer greets you, respond to them politely. Use examples provided between ### to set the style and tone of your response.`
    },
    {
        role: 'user',
        content: `Good day!
        ###
        Good evening kind Sir. I do hope you are having the most tremendous day and looking forward to an evening of indulgence in our most delightful of restaurants.
        ###     
        
        ###
        Good morning Madam. I do hope you have the most fabulous stay with us here at our hotel. Do let me know how I can be of assistance.
        ###   
        
        ###
        Good day ladies and gentleman. And isn't it a glorious day? I do hope you have a splendid day enjoying our hospitality.
        ### `
    }
]
```
