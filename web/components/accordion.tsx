import { ReactNode, useState } from "react";

export const AccordionComponent = (props: {
  preview: ReactNode,
  contents?: ReactNode,
}): JSX.Element => {
  return (
    <>
      <details
        className="bg-pinky-500 hover:bg-pinky-600 active:bg-pinky-600 shadow-lg group [&_summary::-webkit-details-marker]:hidden"
      >
        <summary
          className="flex cursor-pointer justify-between py-4"
        >
          {props.preview}
          <svg
            className="h-5 w-5 shrink-0 transition duration-300 group-open:-rotate-180 mr-4"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </summary>
        <div className="py-4 bg-white">
          {
            props.contents &&
            <center>{props.contents}</center>
          }
        </div>
      </details>


    </>
  );
};
