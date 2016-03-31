class Solution(object):
    def isSelfCrossing(self, x):
        """
        :type x: List[int]
        :rtype: bool
        """
        n = len(x)
        if n < 4:
            return False

        for i in xrange(3, n):
            if x[i] >= x[i-2] and x[i-1] <= x[i-3]:
                return True
            elif i >= 4 and x[i-1] == x[i-3] and x[i]+x[i-4] >= x[i-2]:
                return True
            elif i >= 5 and x[i-2] >= x[i-4] and x[i]+x[i-4] >= x[i-2] and x[i-1] <= x[i-3] and x[i-1]+x[i-5]>=x[i-3]:
                return True
        
        return False
