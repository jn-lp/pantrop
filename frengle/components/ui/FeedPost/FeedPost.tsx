import React from 'react'
import Image from 'next/image'

import {Container} from "@components/ui";
import {Plus, Share, Users} from "@components/icons";

import s from './FeedPost.module.css'
import {ImageLoaderProps} from "next/dist/client/image";

const FeedPost = ({name, username, membersCount, trip, title, startAt, pitStops, tempo}) => {
  const imageLoader = ({src, width}: ImageLoaderProps) => {
    return `https://api.mapbox.com/styles/v1/mapbox/outdoors-v11/static/geojson({"type":"FeatureCollection","features":[{"type":"Feature","properties":{"stroke":"%23F90093","stroke-width":4,"stroke-opacity":0.96},"geometry":{"type":"LineString","coordinates":${src}}}]})/auto/${width}x${~~(width / 1.5)}@2x?padding=${~~(width / 15)}&access_token=pk.eyJ1IjoibGVwZWljbyIsImEiOiJja2JucjlxYnQxdnBzMnRxdnE3bmc1a3JmIn0.6f9cPjWPwmk63Qq_TpuVyg`
  }

  return (
    <article className={s.root}>
      <Container>
        <div className={s.Inner}>
          <div className={s.Header}>
            <div className={s.Group}>
              <Image
                src="/images/me.jpeg"
                alt="Picture of the author"
                width={40}
                height={40}
                className={s.Avatar}
              />
              <div className={s.Name}>
                <h4>{name}</h4>
                <small>@{username}</small>
              </div>
            </div>
            <div className={s.MembersCount}>
              <b>{membersCount}</b>
              <Users/>
            </div>
          </div>
          <Image
            loader={imageLoader}
            src={JSON.stringify(trip)}
            width={300}
            height={200}
            className={s.Map}
          />
          <h3 className={s.Title}>{title}</h3>
          <div className={s.Info}>
            <div>
              <h5>Start</h5>
              <span>
                {
                  startAt
                    .toLocaleString('en-US', {
                      month: 'short',
                      day: 'numeric',
                      hour: 'numeric',
                      minute: 'numeric'
                    })
                }
              </span>
            </div>
            <div>
              <h5>Pit stops</h5>
              <span>{pitStops ? "Yes" : "No"}</span>
            </div>
            <div>
              <h5>Tempo</h5>
              <span>{tempo} km/h</span>
            </div>
          </div>
          <div className={s.Actions}>
            <button>
              <Plus/>
              I'll be
            </button>
            <button>
              <Share/>
              Share
            </button>
          </div>
        </div>
      </Container>
    </article>
  )
}

export default FeedPost
