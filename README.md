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
  </td>
  <td>
  ```json
  {
    "Status" : 200,
    "Message": "OK",
    "Data" : {
        "Title": "How to Draw an Ellipse (Official)",
        "Indonesian": false,
        "Genre": "Korean , Manhwa , Webtoon , Yuri(GL) , Drama , Full Color , Mystery , Office Workers , Romance , Shoujo ai , Thriller , ",
        "MangaUrl": "https://wto.to/series/111564/how-to-draw-an-ellipse-official",
        "ChapterUrl": "https://wto.to/chapter/2925592",
        "ImageUrl": "https://xfs-n12.xfsbb.com/thumb/W300/ampi/4c7/4c72db554a16d59da10ff40e9e8535e5744710e0_1000_1500_486759.jpeg",
        "Id": "2925592",
        "Chapter": "Episode 118",
        "Uploader": "byleth 20 mins ago"
    },...
  }
  ```
  </td>
</table>