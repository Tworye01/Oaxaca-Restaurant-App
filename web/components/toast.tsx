type Props = {
  text: string;
  onClick: () => void;
  onClose: () => void;
  style: "info" | "warn" | "error";
};

export const ToastComponent: React.FC<Props> = ({
  text,
  onClick,
  onClose,
  style
}) => {
  return (
    <div className="flex items-center w-full max-w-xs p-4 mb-4 text-gray-600 border-gray-200 rounded-lg shadow-md hover:border-gray-300 hover:bg-gray-100 hover:shadow-lg">
      <p>{text}</p>
      <button
        onClick={onClose}
        className="ms-auto -mx-1.5 -my-1.5 text-gray-600 hover:text-gray-900 rounded-lg focus:ring-2 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8"
      >
        <svg
          className="w-5 h-5"
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 11.793a1 1 0 1 1-1.414 1.414L10 11.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L8.586 10 6.293 7.707a1 1 0 0 1 1.414-1.414L10 8.586l2.293-2.293a1 1 0 0 1 1.414 1.414L11.414 10l2.293 2.293Z" />
        </svg>
      </button>
    </div>
  );
};
