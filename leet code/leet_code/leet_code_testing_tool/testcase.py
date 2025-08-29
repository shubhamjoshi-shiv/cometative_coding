
def compare_results(result, output_param, input_param, comparison_level):
    match comparison_level:
        case 0:
            if result != output_param:
                print('tc failed','expected ->',output_param,result,'<- got', input_param)
            else:
                print('tc passed')
        case 1:
            if result != output_param:
                print('tc failed','expected ->',output_param,result,'<- got', input_param, sep='\n')
            else:
                print('tc passed')

def run_test_case():
    s = Solution()
    comparison_level = 0

    
    # Example 1
    result = s.maxArea(height = [1,8,6,2,5,4,8,3,7])
    input_param = 'height = [1,8,6,2,5,4,8,3,7]'
    output_param = 49
    compare_results(result, output_param, input_param, comparison_level)
        
    # Example 2
    result = s.maxArea(height = [1,1])
    input_param = 'height = [1,1]'
    output_param = 1
    compare_results(result, output_param, input_param, comparison_level)
        
run_test_case()