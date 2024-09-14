# **CrewInator**

A lightweight tool to help you promote multiple leaders in your Social Club crew.

---

### **How does it work?**

CrewInator works by sending requests in quick succession to promote up to five different people as crew leaders. The trick is in the timing: if the requests are processed before Social Club demotes you, it can promote multiple people while only demoting you, effectively setting up multiple leaders.

### **How to use it?**

1. **Log in** to the account currently holding the crew leader position.
2. **Promote up to 5 people** you want as new leaders to the commissioner rank.
3. **Retrieve your Bearer Token**:
   - Go to your crew’s page.
   - Right-click anywhere on the page and select **Inspect**.
   - Navigate to the **Console** tab.
   - Paste the following JavaScript code into the console and press enter to copy the Bearer Token to your clipboard:
   
   ```javascript
   function gT(){let t=decodeURIComponent(document.cookie).split(";"),e="BearerToken=";for(let o of t)if((o=o.trim()).startsWith(e)){let i=o.substring(e.length,o.length);navigator.clipboard.writeText(i).then(()=>{alert("Copied To Clipboard")}).catch(t=>{console.log("Copy this:",i),alert("Copy this: "+i)});return}alert("You don't seem to be logged in")}gT();
   ```

   > **Note:** This JavaScript code will also be automatically copied to your clipboard when you run the CrewInator program, provided your clipboard is empty. Paste it into your browser's console to extract your Bearer Token.

4. **Run the CrewInator**:

   When you start the program, you'll be prompted to:
   - Enter your **crew name**.
   - Paste the **Bearer Token** (retrieved from the console).

   > **Important:** Make sure the Bearer Token starts with `"eyJ"`, the typical format for a valid Bearer Token.

5. The tool will attempt to promote the people you selected as commissioners in step 2. If the promotions are successful, you'll see a success message in the terminal.

### **Credits**

Special thanks to Nanno's special(ed) internet for discovering this.

### **Notes**

This method is timing-sensitive, so it may not always work. However, using CrewInator should increase the chances of success drastically in comparison to some other poor attempts at doing this. If the promotion fails, check for the following:

- You don’t have leader permissions.
- There’s a typo in your crew name.
- The Rockstar IDs you provided aren’t in the crew.
- Your Bearer Token expired (refresh Social Club and get a new one).
- Your internet is slow/laggy. Consider using https://crew.dudx.info/ instead, as this will use a server instead

---

<small><i>dudx certified</i></small>