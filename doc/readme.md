## App flow
Adding an additional step to clarify the user's preference on how to handle the potentially corrupted files is a great idea. This gives the user more control over the process. The breakdown of the plan with steps:

1. **User Input and Metadata Collection:**
   - User inputs a directory path.
   - You scan metadata (such as file name, size, extension, created/modified dates) for each file in the directory and its subdirectories.
   - You save this metadata in a JSON file.

2. **User Preference for Handling Corrupt Files:**
   - You ask the user if they want to verify for corrupt files.
   - If the user agrees, you proceed with the verification process.
   - You ask the user how they want to handle potentially corrupted files:
     - Option 1: Move corrupted files to a designated "corrupted" folder.
     - Option 2: Generate a list of corrupted files without taking any action.
     - Option 3: Delete the corrupted files.
     - Option 4: Skip corrupt file handling and finish the process.

3. **Hex Signature Comparison:**
   - For each file, you read the initial bytes to obtain the hex signature.
   - You compare the obtained hex signature with the expected hex signature for the corresponding file extension.

4. **Handling Corrupted Files:**
   - Based on the user's preference chosen in step 2, you take the appropriate action:
     - If the user chose option 1, move the corrupted files to a designated "corrupted" folder.
     - If the user chose option 2, generate a list of corrupted files.
     - If the user chose option 3, delete the corrupted files.
     - If the user chose option 4, skip this step and finish the process.

This updated plan provides clear choices for the user, allowing them to decide how they want to handle the potentially corrupted files. It's user-friendly and ensures that the user's preferences are respected throughout the process. As always, testing the plan on a small-scale dataset before full implementation is a good practice to ensure that the user interactions and handling options work as intended.