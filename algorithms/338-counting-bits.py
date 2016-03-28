class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        result = []
        len = 0
        count = 1
        i = 0
        while True:
            if len > num:
                return result
            
            if i >count:
                i = 0
                count = len - 1

            if len == 0 or len == 1:
                result.append(len)
                len += 1
                continue
            
            result.append(1+result[i])
            i += 1
            len +=1
