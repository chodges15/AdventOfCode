# Gemini 3.0

def solve_safe(data):
    # 1. Initialize the dial state
    # The problem states the dial starts at 50
    current_pos = 50 
    password_count = 0
    
    # 2. Process the input line by line
    lines = data.strip().split('\n')
    
    for line in lines:
        line = line.strip()
        if not line:
            continue
            
        # Parse direction (first char) and amount (rest of string)
        direction = line[0]
        amount = int(line[1:])
        
        # 3. Apply Rotation using Modulo Arithmetic
        # The dial has 100 numbers (0-99). 
        # We use % 100 to wrap around the circle automatically.
        if direction == 'R':
            # Right = Add
            current_pos = (current_pos + amount) % 100
        elif direction == 'L':
            # Left = Subtract
            # Python handles negative modulo correctly (e.g., -1 % 100 = 99)
            current_pos = (current_pos - amount) % 100
            
        # 4. Check if the dial is pointing at 0
        if current_pos == 0:
            password_count += 1
            
    return password_count

