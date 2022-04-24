import './index.css';
import liff from '@line/liff'

document.addEventListener("DOMContentLoaded", function() {
  liff
    .init({ liffId: process.env.LIFF_ID })
    .then(() => {
        liff
            .sendMessages([
                
                {
                    "type": "template",
                    "altText": "This is a buttons template",
                    "template": {
                        "type": "buttons",
                        "thumbnailImageUrl": "https://example.com/bot/images/image.jpg",
                        "imageAspectRatio": "rectangle",
                        "imageSize": "cover",
                        "imageBackgroundColor": "#FFFFFF",
                        "title": "Menu",
                        "text": "Please select",
                        "defaultAction": {
                            "type": "uri",
                            "label": "View detail",
                            "uri": "http://example.com/page/123"
                        },
                        "actions": [
                            {
                              "type": "uri",
                              "label": "View detail",
                              "uri": "http://example.com/page/123"
                            }
                        ]
                    }
                  }
            ])
            .then(() => {
                console.log("message sent");
            })
            .catch((err) => {
                console.log("error", err);
            });
        console.log("Success! you can do something with LIFF API here.")
    })
    .catch((error) => {
        console.log(error)
    })
});
