import cx from "classnames";
import React, { useState } from "react";

const LikeDislike = () => {
  const [likes, setLikes] = useState(100);
  const [dislikes, setDislikes] = useState(100);
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);

  const handleLike = () => {
    if (liked) {
      setLiked(false);
      setLikes(likes - 1);
    } else {
      setLikes(likes + 1);
      setLiked(true);

      if (disliked) {
        setDislikes(dislikes - 1);
        setDisliked(false);
      }
    }
  };

  const handleDislike = () => {
    if (disliked) {
      setDisliked(false);
      setDislikes(dislikes - 1);
    } else {
      setDislikes(dislikes + 1);
      setDisliked(true);

      if (liked) {
        setLikes(likes - 1);
        setLiked(false);
      }
    }
  };

  return (
    <>
      <div style={{ display: "flex", gap: "10px" }}>
        <button 
          className={cx("like-button", { liked })} 
          onClick={handleLike}
        >
          Like | <span className="likes-counter">{likes}</span>
        </button>
        <button 
          className={cx("dislike-button", { disliked })}
          onClick={handleDislike}
        >
          Dislike | <span className="dislike-counter">{dislikes}</span>
        </button>
      </div>
      <style>{`
        .like-button, .dislike-button {
            font-size: 1rem;
            padding: 5px 10px;
            color: #585858;
            cursor: pointer;
            border: 1px solid #e0e0e0;
            border-radius: 4px;
            background-color: #f5f5f5;
            transition: all 0.2s;
        }

        .liked {
            font-weight: bold;
            color: #1565c0;
            background-color: #e3f2fd;
            border-color: #1565c0;
        }

        .disliked {
            font-weight: bold;
            color: #c62828;
            background-color: #ffebee;
            border-color: #c62828;
        }
      `}</style>
    </>
  );
};

export default LikeDislike; 