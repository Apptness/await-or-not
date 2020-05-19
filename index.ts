export default async <T>(sync: boolean, callback: () => Promise<T>) => {
  if (sync) {
    return callback();
  } else {
    callback();
  }
};
