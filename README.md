

# **CrewInator**

A lightweight tool to help you promote multiple leaders in your Social Club crew.

---

### **How does it work?**

CrewInator works by sending two requests in quick succession to promote two different people as crew leaders. The trick is in the timing: if the second request is processed before the first request demotes you, Social Club will promote both people while only demoting you, effectively setting up two leaders.

### **How to use it?**

0. **Download** the executable from **out** or [releases](https://github.com/41Baloo/crewInator/releases)
1. **Log in** to the account currently holding the crew leader position.
2. Ensure the **two people** you want to promote as new leaders are members of the crew.
3. **Find your crew's ID**: Visit your crew’s page on Social Club and look for “byname” in the request URL to get the crew ID.
```javascript
// example response from https://scapi.rockstargames.com/crew/byname?name=ign_crew
{
   "crewId":735, // <-- We want this
   "crewName":"IGN Crew",
   "crewTag":"IGN",
   "crewMotto":"Obsessed with Gaming, Entertainment, and Everything guys enjoy.",
   "memberCount":7820294,
   /* ... */
   "isMember":false,
   "canDeleteFromWall":false,
   "createdAt":"0001-01-01T00:00:00.000Z",
   "status":true
}
```
4. **Get your Bearer Token**: Visit your crew’s page and run the following JavaScript in your browser’s console to retrieve your Bearer Token (it expires quickly, so do this step last).

```javascript
// BearerToken script
function gT() {
    let cookies = decodeURIComponent(document.cookie),
        splitCookies = cookies.split(";"),
        tokenPrefix = "BearerToken=";

    for (let cookie of splitCookies) {
        cookie = cookie.trim();
        if (cookie.startsWith(tokenPrefix)) {
            let token = cookie.substring(tokenPrefix.length);
            navigator.clipboard.writeText(token).then(() => {
                alert("Bearer Token copied to clipboard");
            }).catch(() => {
                console.log("Copy this manually:", token);
                alert("Copy this manually: " + token);
            });
            return;
        }
    }

    alert("You don't seem to be logged in.");
}
gT();
```

5. **Run the program** using the following commands:

#### **Linux**
```
./crewInator -crewID=YOUR_CREW_ID -firstRid=FIRST_PERSON_RID -secondRid=SECOND_PERSON_RID -token=YOUR_BEARER_TOKEN
```

#### **Windows**
```
crewInator -crewID=YOUR_CREW_ID -firstRid=FIRST_PERSON_RID -secondRid=SECOND_PERSON_RID -token=YOUR_BEARER_TOKEN
```

### **Credits**

Special thanks to Nanno's special(ed) internet for discovering this.

### **Notes**

This method is timing-sensitive, so it may not always work. However, using CrewInator should increase the chances of success. If the promotion fails, check for the following:

- You don’t have leader permissions.
- There’s a typo in your crew ID.
- The Rockstar IDs you provided aren’t in the crew.
- Your Bearer Token expired (refresh Social Club and get a new one).
- Your internet is slow/laggy

---

<small><i>dudx certified</i></small>