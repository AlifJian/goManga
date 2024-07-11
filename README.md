## How to Run
This API can run localy in your system in port 3000 : 
```
- git clone https://github.com/AlifJian/goManga/
- cd goManga
- go run main.go
```

### MangaScrapper Extension
List of websites available now :
- Mangatoto

## Endpoint
<table>
  <tr>
    <td>Endpoint</td><td>Method</td><td>Status</td><td>Request</td><td>Response</td>
  </tr>
  <tr>
  <td> `/manga` </td>
  <td> GET </td>
  <td> 200 </td>
  <td>

   ```
    Request Params,
    limit : int [Default 10]
    index : int [Defautl 0]
   ```

  </td>
  <td>
 
   ```json
    {
    "Status" : 200,
    "Message": "OK",
    "Data" : [
        {
        "Title": "How to Draw an Ellipse (Official)",
        "Indonesian": false,
        "Genre": "Korean , Manhwa , Webtoon , Yuri(GL) , Drama , Full Color , Mystery , Office Workers , Romance , Shoujo ai , Thriller , ",
        "MangaUrl": "https://wto.to/series/111564/how-to-draw-an-ellipse-official",
        "ChapterUrl": "https://wto.to/chapter/2925592",
        "ImageUrl": "https://xfs-n12.xfsbb.com/thumb/W300/ampi/4c7/4c72db554a16d59da10ff40e9e8535e5744710e0_1000_1500_486759.jpeg",
        "Id": "2925592",
        "Chapter": "Episode 118",
        "Uploader": "byleth 20 mins ago"
    },
    ...
    ]
  }
  ```

  </td>
  </tr>

  <tr>
  <td> `/manga/chapter/:id` </td>
  <td>GET</td>
  <td>200</td>
  <td>   
  </td>
  <td>

   ```json
       {
        "Status" : 200,
        "Message" : "OK",
        "Data" : {
            "imgLength": 62,
            "imgUrl": [
                "https://xfs-n07.xfsbb.com/comic/7006/c34/668f714e1f0782a407a0d43c/58109938_940_1821_44926.webp",
                "https://xfs-n17.xfsbb.com/comic/7006/c34/668f714e1f0782a407a0d43c/58109945_940_1821_36236.webp",
                "https://xfs-n12.xfsbb.com/comic/7006/c34/668f714e1f0782a407a0d43c/58109931_940_1821_75144.webp",
                "https://xfs-n17.xfsbb.com/comic/7006/c34/668f714e1f0782a407a0d43c/58109935_940_1821_23942.webp",
                ...
            ]
        }
    }
   ```
  </td>
  </tr>

  <tr>
  <td> `/manga/search` </td>
  <td>GET</td>
  <td>200</td>
  <td>

   ```
    Request Params,
    title : String [default ""]
    limit : int [Default 10]
    index : int [Defautl 0]
   ```

  </td>
  <td>

   ```json
       {
        "Status" : 200,
        "Message" : "OK",
        "Data" : [
             {
                "Title": "Choujin X [Cryptarithm]",
                "Indonesian": true,
                "Genre": "English , Manga , Seinen(M) , Shounen(B) , ",
                "MangaUrl": "https://wto.to/series/162184/choujin-x-cryptarithm",
                "ChapterUrl": "",
                "ImageUrl": "https://xfs-n02.xfsbb.com/thumb/W300/ampi/9d5/9d59f0acac8e76d8b487422576fe844a49a258ac_600_857_63052.jpeg",
                "Id": "",
                "Chapter": "",
                "Uploader": ""
            },
            ...
        ]
    }
   ```
  </td>
  </tr>
</table>