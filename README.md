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
  <td> `/btoto` </td>
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
        "SeriesId" : "111564",
        "ChapterId": "2925592",
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
  <td> `/btoto/chapter/:id` </td>
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
  <td> `/btoto/search` </td>
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
                "Title": "Before the Spilled Milk Dries",
                "Indonesian": false,
                "Genre": "Japanese , Doujinshi , Yuri(GL) , Romance , Shoujo ai , Tragedy , ",
                "MangaUrl": "https://wto.to/series/117377/before-the-spilled-milk-dries",
                "ChapterUrl": "https://wto.to/chapter/2143116",
                "ImageUrl": "https://xfs-n07.xfsbb.com/thumb/W300/ampi/e8a/e8a5754f166fba5dbcc32c5d71ca488ebfb6c4f2_375_533_83205.jpeg",
                "SeriesId": "117377",
                "ChapterId": "2143116",
                "Chapter": "Ch.4",
                "Uploader": "cherrimorre 558 days ago"
            },
            ...
        ]
    }
   ```
  </td>
  </tr>
</table>