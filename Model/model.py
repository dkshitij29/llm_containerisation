from transformers import LlamaForCausalLM, LlamaTokenizer

# Specify the model name; replace this with the actual Llama model identifier if it's available.
model_name = "facebook/llama"

# Download the model and tokenizer
model = LlamaForCausalLM.from_pretrained(model_name)
tokenizer = LlamaTokenizer.from_pretrained(model_name)

# Save them locally (optional)
model.save_pretrained("/home/dkshitij_debian/Code/model")
tokenizer.save_pretrained("/home/dkshitij_debian/Code/model")
 