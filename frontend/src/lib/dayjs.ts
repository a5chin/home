/* eslint-disable import/no-named-as-default-member */
// eslint-disable-next-line no-restricted-imports
import dayjs from "dayjs";
import "dayjs/locale/ja";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.locale("ja");
dayjs.extend(relativeTime);

export default dayjs;
