import Head from "next/head";
import { useState } from "react";
import { createActivity } from "../service";
import Datetime from 'react-datetime';
import "react-datetime/css/react-datetime.css";

export default function Home(props) {
  /** You can access to liff and liffError object through the props.
   *  const { liff, liffError } = props;
   *  console.log(liff.getVersion());
   *
   *  Learn more about LIFF API documentation (https://developers.line.biz/en/reference/liff)
   **/
  const { liff, liffError } = props;
  const [activity, setActivity] = useState();
  const [date, setDate] = useState(new Date());
  const [place, setPlace] = useState();

  const handleSubmit = async (e) => {
    // Post to golang server
    e.preventDefault()
    try {
      const idToken = liff.getIDToken()
      // await createActivity({ activity, date, place, idToken });

      // Submit message
      
      await liff.sendMessages([
        {
          type: "template",
          altText: "This is a buttons template",
          template: {
            type: "buttons",
            thumbnailImageUrl: "https://example.com/bot/images/image.jpg",
            imageAspectRatio: "rectangle",
            imageSize: "cover",
            imageBackgroundColor: "#FFFFFF",
            title: "Menu",
            text: "Please select",
            defaultAction: {
              type: "uri",
              label: "View detail",
              uri: "http://example.com/page/123",
            },
            actions: [
              {
                type: "postback",
                label: "Attend!",
                data: "action=attend"
              }
            ],
          },
        },
      ]);
      console.log("Success!!")
    } catch (err) {
      console.log("err");
      console.log(err);
    }

    console.log("Success22!!")
  };

  return (
    <div>
      <Head>
        <title>Activity Scheduler</title>
      </Head>
      <div className="home">
        <h1 className="home__title">Fill in your activity!</h1>
        <form onSubmit={handleSubmit} method="post">
          <label id="activity">Activity Name:</label>
          <input
            type="text"
            id="activity"
            name="activity"
            required
            onChange={(e) => setActivity(e.target.value)}
          />
          <br />
          <Datetime />
          <br />
          <label id="name">Place:</label>
          <input
            type="text"
            id="place"
            name="place"
            required
            onChange={(e) => setPlace(e.target.value)}
          />

          <button type="submit">Submit</button>
        </form>
      </div>
    </div>
  );
}
