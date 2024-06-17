import Image from 'next/image';
import React from 'react';

type Props = {
  src: string;
  alt: string;
  width: number;
  height: number;
  objectFit?: 'contain' | 'cover' | 'fill' | 'none' | 'scale-down'; // optional property
};

//functional component that accepts ImageProps
export const ImageComponent: React.FC<Props> = ({
  src,
  alt,
  width,
  height,
  objectFit = 'cover'
}): JSX.Element => {
  return (
    <Image
      src={src}
      alt={alt}
      width={width}
      height={height}
      objectFit={objectFit}
    />
  );
};
